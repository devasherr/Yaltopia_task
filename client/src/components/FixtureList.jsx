import { useNavigate } from "react-router-dom";

export default function FixtureList({ fixtures }) {
  const navigate = useNavigate();

  return (
    <div className="space-y-2 mt-4">
      {fixtures.map((f) => (
        <div
          key={f.id}
          onClick={() => navigate(`/fixture/${f.id}`)}
          className="cursor-pointer border p-2 rounded hover:bg-gray-100"
        >
          {f.homeName} VS {f.awayName}
        </div>
      ))}
    </div>
  );
}
