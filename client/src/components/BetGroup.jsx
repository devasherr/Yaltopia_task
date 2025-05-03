export default function BetGroup({ title, options, children }) {
  return (
    <div className="mb-8">
      <h3 className="text-lg font-bold mb-4 text-gray-200">{title}</h3>
      <div className="grid grid-cols-1 md:grid-cols-2 gap-3">{children}</div>
    </div>
  );
}
