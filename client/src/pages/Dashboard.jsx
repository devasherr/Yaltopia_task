import { useEffect, useState } from "react";
import SportList from "../components/SportList";
import FixtureList from "../components/FixtureList";
import { BASE_URL } from "../utils";
import axios from "axios";

export default function Dashboard() {
  const [selectedSportIndex, setSelectedSportIndex] = useState(null);
  const [sportsData, setSportsData] = useState([
    {
      name: "cricket",
      fixtures: [],
    },
    {
      name: "volleyball",
      fixtures: [],
    },
  ]);

  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const fetchFixtures = async () => {
    try {
      setLoading(true);
      setError(null);

      const cricketResponse = await axios.get(`${BASE_URL}/cricket/fixtures`);
      const volleyballResponse = await axios.get(
        `${BASE_URL}/volleyball/fixtures`
      );

      setSportsData([
        {
          name: "cricket",
          fixtures: cricketResponse.data || [],
        },
        {
          name: "volleyball",
          fixtures: volleyballResponse.data || [],
        },
      ]);
    } catch (err) {
      console.error("Error fetching fixtures:", err);
      setError("Failed to load fixtures.");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchFixtures();
  }, []);

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

            <div className="lg:col-span-3">
              {selectedSportIndex !== null ? (
                <div className="bg-gray-800 rounded-lg border border-gray-700 overflow-hidden">
                  <div className="bg-gradient-to-r from-blue-700 to-blue-900 p-4">
                    <h2 className="text-xl font-bold text-white">
                      {sportsData[selectedSportIndex].name
                        .charAt(0)
                        .toUpperCase() +
                        sportsData[selectedSportIndex].name.slice(1)}{" "}
                      Fixtures
                    </h2>
                  </div>
                  <FixtureList
                    key={selectedSportIndex}
                    fixtures={sportsData[selectedSportIndex].fixtures}
                    sport={sportsData[selectedSportIndex].name}
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
