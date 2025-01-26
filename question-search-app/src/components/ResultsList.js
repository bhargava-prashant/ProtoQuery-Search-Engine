import React from 'react';
import ResultItem from './ResultItem';


function ResultsList({ results, loading, searched }) {
  if (loading) {
    return (
      <div className="loading-message">
        Searching for questions...
      </div>
    );
  }

  if (searched && results.length === 0) {
    return (
      <div className="no-results-message">
        No results found. Try a different query.
      </div>
    );
  }

  return (
    <ul className="results-list">
      {results.map((question) => (
        <ResultItem key={question.ID} question={question} />
      ))}
    </ul>
  );
}

export default ResultsList;