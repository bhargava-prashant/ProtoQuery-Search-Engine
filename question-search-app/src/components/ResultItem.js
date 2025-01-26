import React from 'react';

function ResultItem({ question }) {
  return (
    <li className="results-list-item">
      <div className="result-header">
        {/* <span className="result-id">ID: {question.ID}</span> */}
        <span className="result-type">Type: {question.Type}</span>
      </div>
      <h3 className="result-title">{question.Title}</h3>
    </li>
  );
}

export default ResultItem;