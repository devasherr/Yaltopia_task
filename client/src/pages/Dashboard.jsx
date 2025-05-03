import { useState } from "react";
import { sportsData } from "../data";
import SportList from "../components/SportList";
import FixtureList from "../components/FixtureList";

export default function Dashboard() {
  const [selectedSportIndex, setSelectedSportIndex] = useState(null);

  return (
    <div className="p-6 max-w-3xl mx-auto">
      <h1 className="text-2xl font-bold mb-4">Available Sports</h1>
      <SportList sports={sportsData} onSelect={setSelectedSportIndex} />
      {selectedSportIndex !== null && (
        <>
          <h2 className="text-xl font-medium mt-6">
            Fixtures for {sportsData[selectedSportIndex].name}
          </h2>
          <FixtureList fixtures={sportsData[selectedSportIndex].fixtures} />
        </>
      )}
    </div>
  );
}
