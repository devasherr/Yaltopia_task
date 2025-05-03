package routes

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"yaltopia_task/data"
	"yaltopia_task/models"

	"github.com/gofiber/fiber/v2"
)

func GetCricketFixtures(c *fiber.Ctx) error {
	path := getPath()

	prematchFile, err := os.ReadFile(path + "/cricket_prematch.json")
	if err != nil {
		return err
	}

	resultsFile, err := os.ReadFile(path + "/cricket_result.json")
	if err != nil {
		return err
	}

	prematchData, err := data.ParseCricketPrematchJSON(prematchFile)
	if err != nil {
		return err
	}

	resultsData, err := data.ParseCricketResultJSON(resultsFile)
	if err != nil {
		return err
	}

	var fixtures []models.Fixture

	for _, prematch := range prematchData.Results {
		fixture := models.Fixture{
			ID: prematch.EventID,
		}

		fixtures = append(fixtures, fixture)
	}

	for _, res := range resultsData.Results {
		for i := range fixtures {
			if res.ID == fixtures[i].ID {
				fixtures[i].HomeId = res.Home.ID
				fixtures[i].HomeName = res.Home.Name
				fixtures[i].AwayId = res.Away.ID
				fixtures[i].AwayName = res.Away.Name
			}
		}
	}

	return c.JSON(fixtures)
}

func GetCricket1x2(c *fiber.Ctx) error {
	path := getPath()

	prematchFile, err := os.ReadFile(path + "/cricket_prematch.json")
	if err != nil {
		return err
	}

	prematchData, err := data.ParseCricketPrematchJSON(prematchFile)
	if err != nil {
		return err
	}

	id := c.Params("id")

	var odds []models.Odd

	for _, match := range prematchData.Results {
		if match.EventID == id {
			for i := range match.Main.SP.ToWinTheMatch.Odds {
				odd := models.Odd{
					ID:   match.Main.SP.ToWinTheMatch.Odds[i].ID,
					Odds: match.Main.SP.ToWinTheMatch.Odds[i].Odds,
					Name: match.Main.SP.ToWinTheMatch.Odds[i].Name,
				}

				odds = append(odds, odd)
			}
		}
	}

	return c.JSON(odds)
}

func GetCricketOverUnder(c *fiber.Ctx) error {
	path, err := os.Getwd()

	if err != nil {
		return err
	}

	path = filepath.Join(path, "data/parsed/")
	prematchFile, err := os.ReadFile(path + "/cricket_prematch.json")
	if err != nil {
		return err
	}

	prematchData, err := data.ParseCricketPrematchJSON(prematchFile)
	if err != nil {
		return err
	}

	id := c.Params("id")

	var odds []models.Odd

	for _, prematch := range prematchData.Results {
		if prematch.EventID == id {
			for i := range prematch.Innings1.SP.FirstInningsScore.Odds {
				odd := models.Odd{
					ID:     prematch.Innings1.SP.FirstInningsScore.Odds[i].ID,
					Odds:   prematch.Innings1.SP.FirstInningsScore.Odds[i].Odds,
					Name:   prematch.Innings1.SP.FirstInningsScore.Odds[i].Name,
					Header: prematch.Innings1.SP.FirstInningsScore.Odds[i].Header,
				}

				odds = append(odds, odd)
			}
		}
	}

	return c.JSON(odds)
}

func EvaluateCricket1x2(c *fiber.Ctx) error {
	path := getPath()

	var request models.CricketRequest

	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	resultsFile, err := os.ReadFile(path + "/cricket_result.json")
	if err != nil {
		return err
	}

	results, err := data.ParseCricketResultJSON(resultsFile)
	if err != nil {
		return err
	}

	for _, res := range results.Results {
		if res.ID == request.ID {
			score := strings.Split(res.SS, "-")
			winner := "0"

			x, _ := strconv.Atoi(score[0])
			y, _ := strconv.Atoi(score[1])

			if x > y {
				winner = "1"
			} else if x < y {
				winner = "2"
			}

			verdict := "Wrong"
			if request.Name == winner {
				verdict = "Correct"
			}

			return c.Status(http.StatusOK).JSON(fiber.Map{
				"message": verdict,
			})
		}
	}

	return c.Status(http.StatusNotFound).JSON(fiber.Map{
		"message": "Result for match not found",
	})
}

func EvaluateCricketUnderAbove(c *fiber.Ctx) error {
	path := getPath()

	var request models.CricketRequest

	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	resultsFile, err := os.ReadFile(path + "/cricket_result.json")
	if err != nil {
		return err
	}

	results, err := data.ParseCricketResultJSON(resultsFile)
	if err != nil {
		return err
	}

	for _, res := range results.Results {
		if res.ID == request.ID {
			score := strings.Split(res.SS, "-")

			x, _ := strconv.Atoi(score[0])
			y, _ := strconv.Atoi(score[1])

			totalRuns := x + y
			n, _ := strconv.Atoi(request.Name)

			verdict := "Wrong"

			if request.Header == "Under" && n < totalRuns {
				verdict = "Correct"
			}

			if request.Header == "Over" && n > totalRuns {
				verdict = "Correct"
			}

			return c.Status(http.StatusOK).JSON(fiber.Map{
				"message": verdict,
			})
		}
	}

	return c.Status(http.StatusNotFound).JSON(fiber.Map{
		"message": "Result for match not found",
	})
}

func GetVolleyFixtures(c *fiber.Ctx) error {
	path := getPath()

	prematchFile, err := os.ReadFile(path + "/volleyball_prematch.json")
	if err != nil {
		return err
	}

	resultsFile, err := os.ReadFile(path + "/volleyball_result.json")
	if err != nil {
		return err
	}

	prematchData, err := data.ParseVolleyballPrematchJSON(prematchFile)
	if err != nil {
		return err
	}

	resultsData, err := data.ParseVolleyballResultJSON(resultsFile)
	if err != nil {
		return err
	}

	var fixtures []models.Fixture

	for _, prematch := range prematchData.Results {
		fixture := models.Fixture{
			ID: prematch.EventID,
		}

		fixtures = append(fixtures, fixture)
	}

	// Match with results data to get team names
	for _, res := range resultsData.Results {
		for i := range fixtures {
			if res.ID == fixtures[i].ID {
				fixtures[i].HomeId = res.Home.ID
				fixtures[i].HomeName = res.Home.Name
				fixtures[i].AwayId = res.Away.ID
				fixtures[i].AwayName = res.Away.Name
			}
		}
	}

	return c.JSON(fixtures)
}

func GetVolley1x2(c *fiber.Ctx) error {
	path := getPath()

	prematchFile, err := os.ReadFile(path + "/volleyball_prematch.json")
	if err != nil {
		return err
	}

	prematchData, err := data.ParseVolleyballPrematchJSON(prematchFile)
	if err != nil {
		return err
	}

	id := c.Params("id")

	var odds []models.Odd

	for _, match := range prematchData.Results {
		if match.EventID == id {
			// check if game_lines exists in main SP
			if match.Main.Sp.GameLines.Odds != nil {
				for _, odd := range match.Main.Sp.GameLines.Odds {
					if odd.Header == "1" || odd.Header == "2" {
						odd := models.Odd{
							ID:     odd.ID,
							Odds:   odd.Odds,
							Header: odd.Header,
						}

						odds = append(odds, odd)
					}
				}
			}
		}
	}

	return c.JSON(odds)
}

func GetVolleyOverUnder(c *fiber.Ctx) error {
	path := getPath()

	prematchFile, err := os.ReadFile(path + "/volleyball_prematch.json")
	if err != nil {
		return err
	}

	prematchData, err := data.ParseVolleyballPrematchJSON(prematchFile)
	if err != nil {
		return err
	}

	id := c.Params("id")

	var odds []models.Odd

	for _, match := range prematchData.Results {
		if match.EventID == id {
			// check main SP game_lines for total odds
			if match.Main.Sp.GameLines.Odds != nil {
				for _, odd := range match.Main.Sp.GameLines.Odds {
					if odd.Name == "Total" && (odd.Header == "1" || odd.Header == "2") {
						odds = append(odds, models.Odd{
							ID:       odd.ID,
							Odds:     odd.Odds,
							Name:     odd.Name,
							Header:   odd.Header,
							Handicap: odd.Handicap,
						})
					}
				}
			}

			// check set_1_lines in others if available
			for _, other := range match.Others {
				if other.Sp.Set1Lines.Odds != nil {
					for _, odd := range other.Sp.Set1Lines.Odds {
						if odd.Name == "Total" {
							odds = append(odds, models.Odd{
								ID:       odd.ID,
								Odds:     odd.Odds,
								Name:     odd.Name,
								Header:   odd.Header,
								Handicap: odd.Handicap,
							})
						}
					}
				}
			}
		}
	}

	return c.JSON(odds)
}

func GetVolleyCorrectScore(c *fiber.Ctx) error {
	path := getPath()

	prematchFile, err := os.ReadFile(path + "/volleyball_prematch.json")
	if err != nil {
		return err
	}

	prematchData, err := data.ParseVolleyballPrematchJSON(prematchFile)
	if err != nil {
		return err
	}

	id := c.Params("id")

	var scores []struct {
		Id     string
		Odds   string
		Name   string
		Header string
	}

	for _, match := range prematchData.Results {
		if match.EventID == id {
			if match.Main.Sp.CorrectSetScore.Odds != nil {
				for _, score := range match.Main.Sp.CorrectSetScore.Odds {
					scores = append(scores, struct {
						Id     string
						Odds   string
						Name   string
						Header string
					}{
						Id:     score.ID,
						Odds:   score.Odds,
						Name:   score.Name,
						Header: score.Header,
					})
				}
			}
		}
	}

	return c.JSON(scores)
}

func EvaluateVolley1x2(c *fiber.Ctx) error {
	path := getPath()

	var request models.VolleyballRequest

	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return err
	}

	resultFile, err := os.ReadFile(path + "/volleyball_result.json")
	if err != nil {
		return err
	}

	results, err := data.ParseVolleyballResultJSON(resultFile)
	if err != nil {
		return err
	}

	for _, res := range results.Results {
		if res.ID == request.ID {
			score := strings.Split(res.SS, "-")
			x, _ := strconv.Atoi(score[0])
			y, _ := strconv.Atoi(score[1])

			verdict := "Wrong"
			if request.Header == "1" && x > y {
				verdict = "Correct"
			}
			if request.Header == "2" && x < y {
				verdict = "Correct"
			}
			if request.Header == "0" && x == y {
				verdict = "Correct"
			}

			return c.Status(http.StatusOK).JSON(fiber.Map{
				"message": verdict,
			})
		}
	}

	return c.Status(http.StatusNotFound).JSON(fiber.Map{
		"message": "Result for match not found",
	})
}

func EvaluateVolleyUnderAbove(c *fiber.Ctx) error {
	path := getPath()

	var request models.VolleyballRequest

	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return err
	}

	resultFile, err := os.ReadFile(path + "/volleyball_result.json")
	if err != nil {
		return err
	}

	results, err := data.ParseVolleyballResultJSON(resultFile)
	if err != nil {
		return err
	}

	for _, res := range results.Results {
		if res.ID == request.ID {
			total := 0
			for i := range res.Scores {
				x, _ := strconv.Atoi(res.Scores[i].Home)
				y, _ := strconv.Atoi(res.Scores[i].Away)

				total += (x + y)
			}

			verdict := "Wrong"
			values := strings.Split(request.Handicap, " ")
			v, _ := strconv.Atoi(values[1])

			if values[0] == "O" && v < total {
				verdict = "Correct"
			}
			if values[0] == "U" && v > total {
				verdict = "Correct"
			}

			return c.Status(http.StatusOK).JSON(fiber.Map{
				"message": verdict,
			})
		}
	}

	return c.Status(http.StatusNotFound).JSON(fiber.Map{
		"message": "Result for match not found",
	})
}

func EvaluateVolleyCorrectScore(c *fiber.Ctx) error {
	path := getPath()

	var request models.VolleyballRequest

	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return err
	}

	resultFile, err := os.ReadFile(path + "/volleyball_result.json")
	if err != nil {
		return err
	}

	results, err := data.ParseVolleyballResultJSON(resultFile)
	if err != nil {
		return err
	}

	for _, res := range results.Results {
		if res.ID == request.ID {
			verdict := "Wrong"

			if res.SS == request.Name {
				verdict = "Correct"
			}

			return c.Status(http.StatusOK).JSON(fiber.Map{
				"message": verdict,
			})

		}
	}

	return c.Status(http.StatusNotFound).JSON(fiber.Map{
		"message": "Result for match not found",
	})
}

func getPath() string {
	path, _ := os.Getwd()
	return filepath.Join(path, "data/parsed/")
}
