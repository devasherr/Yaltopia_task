import BetGroup from "../components/BetGroup";
import BetOption from "../components/BetOption";
import { bets } from "../data";

export default function FixturePage() {
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
