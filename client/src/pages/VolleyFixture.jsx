import BetGroup from "../components/BetGroup";
import BetOption from "../components/BetOption";
import { useState, useEffect } from "react";
import { BASE_URL } from "../utils";
import axios from "axios";

export default function FixturePage() {
  const [bets, setBets] = useState({
    "1x2": [],
    underOverVolley: [],
    correctScore: [],
  });

  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const fixtureId = window.location.href.split("/")[5];

  const fetchBets = async () => {
    try {
      setLoading(true);
      setError(null);

      const volley_1x2 = await axios.get(
        `${BASE_URL}/volleyball/1x2/${fixtureId}`
      );
      const volley_overunder = await axios.get(
        `${BASE_URL}/volleyball/over-under/${fixtureId}`
      );

      const volley_correctscore = await axios.get(
        `${BASE_URL}/volleyball/correct-score/${fixtureId}`
      );

      setBets({
        "1x2": volley_1x2.data || [],
        underOverVolley: volley_overunder.data || [],
        correctScore: volley_correctscore.data || [],
      });
    } catch (err) {
      console.error("Error fetching bets:", err);
      setError("Failed to load betting options.");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchBets();
  }, [fixtureId]);

  if (loading) {
    return (
      <div className="flex justify-center items-center h-screen bg-gray-900">
        <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-blue-500"></div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex justify-center items-center h-screen bg-gray-900 text-white">
        {error}
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-900 p-6 mx-auto">
      <h1 className="text-2xl font-bold mb-6 text-white">
        Fixture Betting Options
      </h1>

      <BetGroup title="Match Result (1X2)">
        {bets["1x2"].map((opt) => (
          <BetOption
            key={opt.id}
            label={
              opt.header === "1"
                ? "Home Win"
                : opt.header === "2"
                ? "Away Win"
                : "Draw"
            }
            odds={opt.odds}
          />
        ))}
      </BetGroup>

      <BetGroup title="Points Over/Under">
        {bets["underOverVolley"].map((opt) => (
          <BetOption
            key={opt.id}
            label={`Team ${opt.header}`}
            odds={opt.odds}
            sublabel={opt.handicap}
          />
        ))}
      </BetGroup>

      <BetGroup title="Correct Score (Sets)">
        <div className="md:col-span-2">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <h4 className="text-md font-semibold mb-2 text-gray-300">
                Home Team Wins
              </h4>
              <div className="space-y-2">
                {bets["correctScore"]
                  .filter((opt) => opt.Header === "1")
                  .map((opt) => (
                    <BetOption key={opt.Id} label={opt.Name} odds={opt.Odds} />
                  ))}
              </div>
            </div>
            <div>
              <h4 className="text-md font-semibold mb-2 text-gray-300">
                Away Team Wins
              </h4>
              <div className="space-y-2">
                {bets["correctScore"]
                  .filter((opt) => opt.Header === "2")
                  .map((opt) => (
                    <BetOption key={opt.Id} label={opt.Name} odds={opt.Odds} />
                  ))}
              </div>
            </div>
          </div>
        </div>
      </BetGroup>
    </div>
  );
}
