export default function BetOption({ label, odds, sublabel }) {
  return (
    <div className="border border-gray-700 rounded-lg px-4 py-3 flex justify-between items-center hover:bg-gray-700 cursor-pointer transition-colors">
      <div>
        <span className="font-medium text-gray-100">{label}</span>
        {sublabel && (
          <div className="text-xs text-gray-400 mt-1">{sublabel}</div>
        )}
      </div>
      <span className="bg-blue-900 bg-opacity-40 text-blue-400 font-bold px-3 py-1 rounded">
        {odds}
      </span>
    </div>
  );
}
