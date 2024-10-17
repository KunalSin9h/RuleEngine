import { useState, useEffect } from "react";

type Rule = {
  Name: string;
  Description: string;
  Rule: string,
}

export default function TestRule() {
  const [rules, setRules] = useState<Rule[]>([]);
  const [error, setError] = useState(false);
  const [showModal, setShowModal] = useState(false);
  const [selectedRule, setSelectedRule] = useState<Rule>();

  useEffect(() => {
    const fetchRules = async () => {
      try {
        const response = await fetch('/rules');
        
        if (response.status !== 200) {
          setError(true);
          console.error('Failed to fetch rules:', response.status);
        } else {
          const data = await response.json();
          setRules(data);
        }
      } catch (error) {
        console.error('Error fetching rules:', error);
      }
    };

    fetchRules();
  }, []);

  
  return (
    <div>
      <h1 className="text-white font-bold">All Rules</h1>
      {error && <p className="text-red-500">Error fetching rules. Oops, something went wrong!</p>}
      <ul className="flex flex-col gap-4">
        {rules.map((rule, index) => (
          <button 
            key={index} 
            className="bg-gray-700 text-white px-4 py-2 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 cursor-pointer hover:bg-gray-600 flex flex-col"
            onClick={() => {
              setSelectedRule(rule);
              setShowModal(true);
            }}
          >
            {/*  @ts-ignore*/}
            <span className="font-semibold">{rule.Name}</span>
            {/*  @ts-ignore*/}
            <span className="text-xs text-gray-400">{rule.Description}</span>
          </button>
        ))}
      </ul>
      {showModal && (
        <div className="fixed inset-0 flex items-center justify-center z-50 bg-black bg-opacity-50">
          <div className="flex flex-col gap-4 bg-white rounded-lg shadow-lg p-6 w-2/3">
            <div>
                <h2 className="text-xl font-bold mb-1">Test Rule</h2>
                {/* @ts-ingore */}
                <h2 className="text-lg font-bold text-blue-500 mb-1">{selectedRule?.Name}</h2>
                <p className="text-gray-400 text-xs">{selectedRule?.Description}</p>
            </div>
            <div className="grid grid-cols-2 gap-4 mb-4">
              <div className="w-full">
                <pre className="bg-black text-white p-4 rounded-md  overflow-y-auto">{selectedRule?.Rule}</pre>
              </div>
              <div>
                <textarea

                  className="bg-gray-700 text-white px-4 py-2 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 resize-none w-full"
                  placeholder="Enter user data"
                  style={{ height: '80%' }}
                />
                <button className="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500">Test</button>
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