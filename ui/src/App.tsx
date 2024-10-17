import AddRule from "./AddRule";
import { useState } from "react";
import TestRule from "./TestRule";
import CombineRule from "./CombineRule";

function App() {
  const [showAddRule, setShowAddRule] = useState(true);
  const [showTestRule, setShowTestRule] = useState(false);
  const [showCombineRule, setShowCombineRule] = useState(false);

  return (
    <div className="flex flex-col items-center justify-center bg-gray-800 min-h-screen">
      <div className="gap-4 flex items-center">
      <button
        className={`bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mb-4 ${showAddRule ? 'bg-blue-700' : ''}`}
        onClick={() => {
          setShowAddRule(true);
          setShowTestRule(false);
          setShowCombineRule(false);
        }}
      >
        Add Rule
      </button>
      <button
        className={`bg-yellow-500 hover:bg-yellow-700 text-white font-bold py-2 px-4 rounded mb-4 ${showCombineRule ? 'bg-yellow-700' : ''}`}
        onClick={() => {
          setShowCombineRule(true);
          setShowAddRule(false);
          setShowTestRule(false);
        }}
      >
        Combine Rule
      </button>
      <button
        className={`bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded mb-4 ${showTestRule ? 'bg-green-700' : ''}`}
        onClick={() => {
          setShowTestRule(true);
          setShowAddRule(false);
          setShowCombineRule(false);
        }}
      >
        Test Rule
      </button>

      </div>
      <div className="card p-4 w-full h-full flex items-center justify-center">
        {showAddRule && <AddRule />}
        {showCombineRule && <CombineRule />}
        {showTestRule && <TestRule />}
      </div>
    </div>
  );
}

export default App;
