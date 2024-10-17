import { useState } from "react";

export default function TestRule() {
  const [rule, setRule] = useState('');

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setRule(event.target.value);
  };

  const handleAddRule = () => {
    // Implement logic to add the rule
    console.log('Adding rule:', rule);
    setRule('');
  };
  return (
    <div className="flex flex-col gap-4 w-full">
      <label htmlFor="rule" className="text-white font-bold">Rule:</label>
      <textarea
        id="rule"
        value={rule}
        // @ts-expect-error
        onChange={handleInputChange}
        className="bg-gray-700 text-white px-4 py-2 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 resize-none w-full"
        style={{ width: '100%', height: "200px" }}
      />
      <button
        onClick={handleAddRule}
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 flex items-center justify-center"
      >
        Add
      </button>
    </div>
  );
}