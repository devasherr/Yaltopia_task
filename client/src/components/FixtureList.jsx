import { useNavigate } from "react-router-dom";

export default function FixtureList({ fixtures, sport }) {
  const navigate = useNavigate();

  if (!fixtures || fixtures.length === 0) {
    return (
      <div className="p-8 text-center text-gray-400">
        No fixtures available for this sport
      </div>
    );
  }

  return (
    <div className="divide-y divide-gray-700">
      {fixtures.map((f) => (
        <div
          key={f.id}
          onClick={() => navigate(`/fixture/${sport}/${f.id}`)}
          className="cursor-pointer hover:bg-gray-700 transition-colors duration-150 p-4"
        >
          <div className="flex items-center justify-between">
            <div className="flex-1 flex items-center justify-center space-x-4">
              <div className="flex items-center space-x-2">
                <div className="w-8 h-8 bg-red-900 bg-opacity-40 rounded-full flex items-center justify-center">
                  <span className="text-red-400 font-bold text-xs">H</span>
                </div>
                <span className="font-medium text-gray-200">{f.homeName}</span>
              </div>

              <div className="px-2 py-1 bg-gray-600 rounded-md text-xs font-bold text-gray-300">
                VS
              </div>

              <div className="flex items-center space-x-2">
                <span className="font-medium text-gray-200">{f.awayName}</span>
                <div className="w-8 h-8 bg-blue-900 bg-opacity-40 rounded-full flex items-center justify-center">
                  <span className="text-blue-400 font-bold text-xs">A</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      ))}
    </div>
  );
}
