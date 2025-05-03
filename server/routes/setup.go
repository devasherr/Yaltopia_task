package routes

import "github.com/gofiber/fiber/v2"

func RegisterCricketRoutes(router fiber.Router) {
	router.Get("/fixtures", GetCricketFixtures)
	router.Get("/1x2/:id", GetCricket1x2)
	router.Get("/over-under/:id", GetCricketOverUnder)
	router.Post("/evaluate-1x2", EvaluateCricket1x2)
	router.Post("/evaluate-under-above", EvaluateCricketUnderAbove)
}

func RegisterVolleyballRoutes(router fiber.Router) {
	router.Get("/fixtures", GetVolleyFixtures)
	router.Get("/1x2/:id", GetVolley1x2)
	router.Get("/over-under/:id", GetVolleyOverUnder)
	router.Get("/correct-score/:id", GetVolleyCorrectScore)
	router.Post("/evaluate-1x2", EvaluateVolley1x2)
	router.Post("/evaluate-under-above", EvaluateVolleyUnderAbove)
	router.Post("/evaluate-correct-score", EvaluateVolleyCorrectScore)
}
