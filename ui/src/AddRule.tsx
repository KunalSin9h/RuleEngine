import { useState } from "react";

export default function AddRule() {
  const [rule, setRule] = useState('');
  const [name, setName] = useState('');
  const [description, setDescription] = useState('');
  const [error, setError] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [apiError, setApiError] = useState("");
  const [showModal, setShowModal] = useState(true);
  const [data, setData] = useState('{"age": 30}');

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setRule(event.target.value);
  };

  const handleNameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setName(event.target.value);
  };

  const handleDescriptionChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setDescription(event.target.value);
  }

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (rule.trim() === '' || name.trim() === '') {
      setError(true);
      return;
    }
    setIsLoading(true);
   
    const response = await fetch('/rule', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ rule, name, description }),
    });

    if (response.status !== 200) {
      setApiError(`Failed, got ${response.status}`);
      setIsLoading(false);
      return;
    }

    const data = await response.json();
    setData(data);
    
    setRule('');
    setName('');
    setDescription('');

    setError(false);
    setIsLoading(false);
    setApiError("");
    setShowModal(true);
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
      {error && <p className="text-red-500">Rule and Name are required fields</p>}
      {apiError && <p className="text-red-500">API Error: {apiError}</p>}
      <div className="flex gap-4">
        <div>
          <label htmlFor="name" className="text-white font-bold">Name:</label>
          <input
            id="name"
            value={name}
            onChange={handleNameChange}
            className="bg-gray-700 text-white px-4 py-2 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 resize-none w-full"
          />
        </div>
        <div>
          <label htmlFor="description" className="text-white font-bold">Description:</label>
          <input
            id="description"
            value={description}
            onChange={handleDescriptionChange}
            className="bg-gray-700 text-white px-4 py-2 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 resize-none w-full"
          />
        </div>
      </div>
      <button
        // @ts-expect-error
        onClick={handleSubmit}
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 flex items-center justify-center"
        disabled={isLoading}
      >
        {isLoading ? <span className="animate-spin rounded-full h-5 w-5 border-t-2 border-b-2 border-white" /> : 'Add'}
      </button>
      {showModal && (
        <div className="fixed inset-0 flex items-center justify-center z-50 bg-black bg-opacity-50">
          <div className="flex flex-col gap-4 bg-white rounded-lg shadow-lg p-6 w-96">
            <div>
                <h2 className="text-xl font-bold mb-1">Rule Added Successfully</h2>
                <p className="">The rule has been added successfully.</p>
            </div>
            {data && < pre className="bg-black text-white p-4 rounded-md" >{data}</pre>}
            <button onClick={() => setShowModal(false)} className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500">Close</button>
          </div>
        </div>
      )}
    </div>
  );
}