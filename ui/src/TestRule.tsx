import { useState, useEffect } from "react";

export default function TestRule() {
  const [rules, setRules] = useState([]);
  const [error, setError] = useState(false);

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
          <li key={index} className="bg-gray-700 text-white px-4 py-2 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 cursor-pointer hover:bg-gray-600 flex flex-col">
            {/*  @ts-ignore*/}
            <span className="font-semibold">{rule.Name}</span>
            {/*  @ts-ignore*/}
            <span className="text-xs text-gray-400">{rule.Description}</span>
          </li>
        ))}
      </ul>
    </div>
  );
}