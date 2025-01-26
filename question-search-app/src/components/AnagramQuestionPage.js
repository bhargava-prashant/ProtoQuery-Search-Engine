import React, { useState } from 'react';

function AnagramQuestionPage({ question, onBack }) {
  const [selectedBlocks, setSelectedBlocks] = useState([]);
  const [solved, setSolved] = useState(false);

  const handleBlockClick = (block) => {
    if (solved) return;

    const newSelectedBlocks = [...selectedBlocks];
    
    if (newSelectedBlocks.includes(block)) {
      const index = newSelectedBlocks.indexOf(block);
      newSelectedBlocks.splice(index, 1);
    } else {
      newSelectedBlocks.push(block);
    }

    setSelectedBlocks(newSelectedBlocks);

    // Check if solution matches
    const currentSolution = newSelectedBlocks.map(b => b.text).join('');
    if (currentSolution === question.solution) {
      setSolved(true);
    }
  };

  return (
    <div className="anagram-question-container p-6 bg-white rounded-lg shadow-md">
      <button 
        onClick={onBack} 
        className="mb-4 px-4 py-2 bg-gray-200 rounded hover:bg-gray-300 transition-colors"
      >
        ← Back to List
      </button>
      
      <h2 className="text-xl font-bold mb-4">{question.Title}</h2>
      
      <div className="blocks-container flex flex-wrap gap-2 mb-4">
        {question.blocks.filter(block => block.showInOption).map((block, index) => (
          <div 
            key={index}
            onClick={() => handleBlockClick(block)}
            className={`
              block-item 
              p-2 
              rounded 
              cursor-pointer 
              transition-colors 
              ${selectedBlocks.includes(block) ? 'bg-blue-200' : 'bg-gray-100'}
              ${solved && block.isAnswer ? 'bg-green-200' : ''}
            `}
          >
            {block.text}
          </div>
        ))}
      </div>

      <div className="selected-solution mb-4">
        <strong>Your Solution:</strong> {selectedBlocks.map(b => b.text).join('')}
      </div>

      {solved && (
        <div className="success-message text-green-600 font-bold">
          Congratulations! You solved the anagram.
        </div>
      )}
    </div>
  );
}

export default AnagramQuestionPage;