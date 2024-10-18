import { useEffect, useState } from "react";

export default function CombineRule() {
  const [rules, setRules] = useState<string[]>([]);
  const [name, setName] = useState('');
  const [description, setDescription] = useState('');
  const [isValid, setIsValid] = useState(false);
  const [nameError, setNameError] = useState(false);
  const [data, setData] = useState<any>(null);
  const [showModal, setShowModal] = useState(false);

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

  useEffect(() => {
    setIsValid(name.trim() !== '');
    setNameError(name.trim() === '');
  }, [name]);

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
            onKeyDown={(e) => {
              if (e.key === ' ') {
                e.preventDefault();
              }
            }}
            onChange={(e) => setName(e.target.value)}
            className="bg-gray-700 text-white px-4 py-2 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 resize-none w-full"
            placeholder="Name"
            required
          />
          {nameError && <p className="text-red-500">Name is required</p>}
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
        disabled={!isValid}

        onClick={async (e) => {
          e.preventDefault();
          try {
            const response = await fetch('/rules', {
              method: 'POST',
              headers: {
                'Content-Type': 'application/json',
              },
              body: JSON.stringify({
                rules,
                name,
                description,
              }),
            });

            if (response.status !== 200) {
              // failed
            } else {
              const data = await response.json();
              setData(data);
              setShowModal(true);
            }
          } catch (error) {
            console.error('Error adding rules:', error);
          }
        }}
      >
        Add Rules
      </button>
      {showModal && (
        <div className="fixed inset-0 flex items-center justify-center z-50 bg-black bg-opacity-50">
          <div className="flex flex-col gap-4 bg-white rounded-lg shadow-lg p-6 w-2/3">
            <div className="flex items-center justify-between">
                <h2 className="text-xl font-bold mb-1">Combined Rule</h2>
                <p className="text-sm font-bold text-green-400">Added: {name}</p>
            </div>
            <div className="mb-4">
              <div className="w-full">
                <pre className="bg-black text-white p-4 rounded-md  overflow-y-auto">{JSON.stringify(data, null, 2)}</pre>
              </div>
            </div>
            <div className="flex flex-col gap-2">
              <button onClick={() => setShowModal(false)} className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">Close</button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}