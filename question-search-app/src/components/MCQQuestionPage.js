import React, { useState } from 'react';

function MCQQuestionPage({ question, onBack }) {
  const [selectedAnswer, setSelectedAnswer] = useState(null);
  const [showResult, setShowResult] = useState(false);

  const handleOptionClick = (option) => {
    setSelectedAnswer(option);
    setShowResult(true);
  };

  return (
    <div className="mcq-question-container p-6 bg-white rounded-lg shadow-md">
      <button 
        onClick={onBack} 
        className="mb-4 px-4 py-2 bg-gray-200 rounded hover:bg-gray-300 transition-colors"
      >
        ← Back to List
      </button>
      
      <h2 className="text-xl font-bold mb-4">{question.Title}</h2>
      
      <div className="options-container">
        {question.options.map((option, index) => (
          <div 
            key={index}
            onClick={() => handleOptionClick(option)}
            className={`
              option-item 
              p-3 
              mb-2 
              rounded 
              cursor-pointer 
              transition-colors 
              ${showResult && option.isCorrectAnswer ? 'bg-green-200' : ''}
              ${showResult && selectedAnswer === option && !option.isCorrectAnswer ? 'bg-red-200' : ''}
              ${!showResult ? 'hover:bg-blue-100 bg-gray-100' : ''}
            `}
          >
            {option.text}
            {showResult && option.isCorrectAnswer && (
              <span className="ml-2 text-green-600 font-bold">✓ Correct Answer</span>
            )}
            {showResult && selectedAnswer === option && !option.isCorrectAnswer && (
              <span className="ml-2 text-red-600 font-bold">✗ Incorrect</span>
            )}
          </div>
        ))}
      </div>
    </div>
  );
}

export default MCQQuestionPage;