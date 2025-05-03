export default function SportList({ sports, onSelect }) {
  return (
    <div className="space-y-4">
      {sports.map((sport, index) => (
        <div
          key={index}
          className="cursor-pointer text-lg font-semibold hover:text-blue-600"
          onClick={() => onSelect(index)}
        >
          {sport.name}
        </div>
      ))}
    </div>
  );
}
