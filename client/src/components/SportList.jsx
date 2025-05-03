export default function SportList({ sports, onSelect }) {
  return (
    <div className="space-y-2">
      {sports.map((sport, index) => (
        <div
          key={index}
          className="cursor-pointer p-3 rounded-md hover:bg-gray-600 transition-colors duration-200 flex items-center"
          onClick={() => onSelect(index)}
        >
          <div className="w-8 h-8 mr-3 bg-blue-900 bg-opacity-40 rounded-full flex items-center justify-center">
            <span className="text-blue-400 text-xs font-bold">
              {sport.name.charAt(0)}
            </span>
          </div>
          <span className="font-medium text-gray-200 hover:text-blue-400">
            {sport.name}
          </span>
          <span className="ml-auto bg-gray-600 text-gray-300 text-xs font-bold px-2 py-1 rounded">
            {sport.fixtures.length}
          </span>
        </div>
      ))}
    </div>
  );
}
