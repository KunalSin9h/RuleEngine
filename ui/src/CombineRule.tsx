import { useEffect, useState } from "react";

export default function CombineRule() {
  const [rules, setRules] = useState<string[]>([]);
  const [name, setName] = useState('');
  const [description, setDescription] = useState('');

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>, index: number) => {
    const updatedRules = [...rules];
    updatedRules[index] = event.target.value;
    setRules(updatedRules);
  };

  const handleAddRule = () => {
    setRules([...rules, '']);
  };

  const handleRemoveRule = (index: number) => {
    const updatedRules = [...rules];
    updatedRules.splice(index, 1);
    setRules(updatedRules);
  };

  useEffect(()=> {
    handleAddRule();
  }, [])

  return (
    <div className="flex flex-col gap-4 w-2/4">
      <label htmlFor="rule" className="text-white font-bold">Rules:</label>
      {rules.map((rule, index) => (
        <div key={index} className="flex flex-col gap-2">
          <textarea
            id={`rule-${index}`}
            value={rule}
            // @ts-expect-error
            onChange={(e) => handleInputChange(e, index)}
            className="bg-gray-700 text-white px-4 py-2 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 resize-none w-full"
          />
          <div className="flex items-center gap-4">
          <button
            onClick={() => handleRemoveRule(index)}
            className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded-md focus:outline-none focus:ring-2 focus:ring-red-500 flex items-center justify-center"
          >
            Remove
          </button>
          {index == rules.length - 1 && (
            <button
              onClick={handleAddRule}
              className="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500 flex items-center justify-center"
            >
              Add
            </button>
          )}
        </div>
        </div>
      ))}
      <div className="flex gap-4">
          <input
            id="name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            className="bg-gray-700 text-white px-4 py-2 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 resize-none w-full"
            placeholder="Name"
          />
          <textarea
            id="description"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            className="bg-gray-700 text-white px-4 py-2 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 resize-none w-full"
            placeholder="Description"
          />
      </div>
      <button
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 flex items-center justify-center"


        onClick={(e) => {
          e.preventDefault();
        }}
      >
        Add Rules
      </button>
    </div>
  );
}