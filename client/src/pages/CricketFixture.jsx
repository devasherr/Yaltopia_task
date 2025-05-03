import BetGroup from "../components/BetGroup";
import BetOption from "../components/BetOption";
import { bets } from "../data";
import { useState } from "react";

export default function FixturePage() {
  const [sport, setSport] = useState("cricket");

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
