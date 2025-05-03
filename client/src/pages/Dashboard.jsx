import { useState } from "react";
import { sportsData } from "../data";
import SportList from "../components/SportList";
import FixtureList from "../components/FixtureList";

export default function Dashboard() {
  const [selectedSportIndex, setSelectedSportIndex] = useState(null);

  return (
    <div className="min-h-screen bg-gray-900 p-4 md:p-8">
      <div className="max-w-6xl mx-auto bg-gray-800 rounded-xl shadow-lg overflow-hidden">
        <div className="p-6 md:p-8">
          <div className="flex items-center justify-between mb-6">
            <h1 className="text-2xl md:text-3xl font-bold text-white">
              Sports Betting
            </h1>
            <div className="bg-green-500 text-gray-900 px-4 py-1 rounded-full text-sm font-semibold">
              LIVE
            </div>
          </div>

          <div className="grid grid-cols-1 lg:grid-cols-4 gap-6">
            {/* Sports Sidebar */}
            <div className="lg:col-span-1 bg-gray-700 rounded-lg p-4 border border-gray-600">
              <h2 className="text-lg font-bold text-gray-300 mb-4 pb-2 border-b border-gray-600">
                Available Sports
              </h2>
              <SportList sports={sportsData} onSelect={setSelectedSportIndex} />
            </div>

            {/* Main Content */}
            <div className="lg:col-span-3">
              {selectedSportIndex !== null ? (
                <div className="bg-gray-800 rounded-lg border border-gray-700 overflow-hidden">
                  <div className="bg-gradient-to-r from-blue-700 to-blue-900 p-4">
                    <h2 className="text-xl font-bold text-white">
                      {sportsData[selectedSportIndex].name} Fixtures
                    </h2>
                  </div>
                  <FixtureList
                    key={selectedSportIndex}
                    fixtures={sportsData[selectedSportIndex].fixtures}
                  />
                </div>
              ) : (
                <div className="bg-gray-700 border border-gray-600 rounded-lg p-8 text-center">
                  <svg
                    className="w-12 h-12 mx-auto text-blue-400"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M13 10V3L4 14h7v7l9-11h-7z"
                    />
                  </svg>
                  <h3 className="mt-4 text-lg font-medium text-blue-300">
                    Select a sport to view fixtures
                  </h3>
                  <p className="mt-2 text-blue-400">
                    Choose from the available sports to see upcoming matches
                  </p>
                </div>
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
