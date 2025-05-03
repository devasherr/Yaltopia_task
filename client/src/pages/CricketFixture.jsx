import BetGroup from "../components/BetGroup";
import BetOption from "../components/BetOption";
import { useEffect, useState } from "react";
import { BASE_URL } from "../utils";
import axios from "axios";

export default function FixturePage() {
  const [bets, setBets] = useState({
    "1x2": [],
    underOverCricket: [],
  });
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const fixtureId = window.location.href.split("/")[5];

  const fetchBets = async () => {
    try {
      setLoading(true);
      setError(null);

      const cricket_1x2 = await axios.get(
        `${BASE_URL}/cricket/1x2/${fixtureId}`
      );
      const cricket_overunder = await axios.get(
        `${BASE_URL}/cricket/1x2/${fixtureId}`
      );

      setBets({
        "1x2": cricket_1x2.data || [],
        underOverCricket: cricket_overunder.data || [],
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
              opt.name === "1"
                ? "Home Win"
                : opt.name === "2"
                ? "Away Win"
                : "Draw"
            }
            odds={opt.odds}
          />
        ))}
      </BetGroup>

      <BetGroup title="Runs Over/Under">
        {bets["underOverCricket"].map((opt) => (
          <BetOption
            key={opt.id}
            label={`${opt.header} ${opt.name}`}
            odds={opt.odds}
            sublabel="Total Runs"
          />
        ))}
      </BetGroup>
    </div>
  );
}
