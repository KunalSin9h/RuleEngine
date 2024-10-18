import { useState, useEffect } from "react";

type Rule = {
  Name: string;
  Description: string;
  Rule: string,
  ID: number,
}

export default function TestRule() {
  const [rules, setRules] = useState<Rule[]>([]);
  const [user, setUser] = useState("");
  const [showModal, setShowModal] = useState(false);
  const [selectedRule, setSelectedRule] = useState<Rule>();
  const [gotResult, setGotResult] = useState(false);
  const [result, setResult] = useState(false);
  const [fetchError, setFetchError] = useState(false);

  useEffect(() => {
    const fetchRules = async () => {
      try {
        const response = await fetch('/rules');
        
        if (response.status !== 200) {
          setFetchError(true);
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
      {fetchError && <p className="text-red-500">Error fetching rules. Please try again later.</p>}
      <ul className="flex flex-col gap-4">
        {rules && rules.map((rule, index) => (
          <button 
            key={index} 
            className="bg-gray-700 text-white px-4 py-2 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 cursor-pointer hover:bg-gray-600 flex flex-col"
            onClick={() => {
              setSelectedRule(rule);
              setShowModal(true);
            }}
          >
          </button>
        ))}
        {!rules && <p className="text-white font-mono mt-4">No Rules</p>}
      </ul>
      {showModal && (
        <div className="fixed inset-0 flex items-center justify-center z-50 bg-black bg-opacity-50">
          <div className="flex flex-col gap-4 bg-white rounded-lg shadow-lg p-6 w-2/3">
            <div className="flex items-center justify-between">
              <div>
                <h2 className="text-xl font-bold mb-1">Test Rule</h2>
                {/* @ts-ingore */}
                <h2 className="text-lg font-bold text-blue-500 mb-1">{selectedRule?.Name}</h2>
                <p className="text-gray-400 text-xs">{selectedRule?.Description}</p>
                </div>
                <p className="text-lg font-bold mt-2">
                  {gotResult && (
                    <span className={result ? "text-green-500" : "text-red-500"}>
                      {result ? "🎉 Passed!" : "😭 Failed!"}
                    </span>
                  )}
                </p>
            </div>
            <div className="grid grid-cols-2 gap-4 mb-4">
              <div className="w-full">
                <pre className="bg-black text-white p-4 rounded-md  overflow-y-auto">{selectedRule?.Rule}</pre>
              </div>
              <div>
                <textarea
                  value={user}
                  onChange={(e) => setUser(e.target.value)}
                  className="bg-gray-700 text-white px-4 py-2 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 resize-none w-full"
                  placeholder="Enter user data"
                  style={{ height: '80%' }}
                />
                <button className="bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500"

                  onClick={async (e) => {
                    e.preventDefault();

                    const response = await fetch('/rule/eval', {
                      method: "POST",
                      headers: {
                        'Content-Type': 'application/json',
                     },
                      body: JSON.stringify({
                        rule_id: selectedRule?.ID,
                        user: user,
                      }),
                    });
                 
                    if (response.status !== 200) {
                      setFetchError(true);
                    }else {
                      const data = await response.json();
                      setGotResult(true);
                      setResult(data.result)
                    }
                  }}

                >Test</button>
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