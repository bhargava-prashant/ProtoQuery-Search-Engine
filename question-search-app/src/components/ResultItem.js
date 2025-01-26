import React, { useState } from 'react';
import MCQQuestionPage from './MCQQuestionPage';
import AnagramQuestionPage from './AnagramQuestionPage';

function ResultItem({ question }) {
  const [isDetailView, setIsDetailView] = useState(false);

  const handleItemClick = () => {
    setIsDetailView(true);
  };

  const handleBackToList = () => {
    setIsDetailView(false);
  };

  if (isDetailView) {
    switch (question.Type) {
      case 'MCQ':
        return <MCQQuestionPage question={question} onBack={handleBackToList} />;
      case 'ANAGRAM':
        return <AnagramQuestionPage question={question} onBack={handleBackToList} />;
      default:
        return null;
    }
  }

  return (
    <li 
      className="results-list-item cursor-pointer hover:bg-gray-100 p-4 rounded-lg transition-colors"
      onClick={handleItemClick}
    >
      <div className="result-header">
        <span className="result-type text-sm text-gray-600">Type: {question.Type}</span>
      </div>
      <h3 className="result-title font-semibold text-lg">{question.Title}</h3>
      {question.Type === 'MCQ' && (
        <span className="text-sm text-blue-600 hover:underline">Click to solve MCQ</span>
      )}
      {question.Type === 'ANAGRAM' && (
        <span className="text-sm text-green-600 hover:underline">Click to solve Anagram</span>
      )}
    </li>
  );
}

export default ResultItem;