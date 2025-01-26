import React from 'react';

const Filter = ({ filterType, setFilterType }) => {
  return (
    <div className="mb-4 flex items-center">
      <label htmlFor="question-type" className="mr-2 text-lg">Filter by Question Type:</label>
      <select
        id="question-type"
        value={filterType}
        onChange={(e) => setFilterType(e.target.value)}
        className="px-4 py-2 border rounded-md text-lg"
      >
        <option value="">All</option>
        <option value="MCQ">MCQ</option>
        <option value="ANAGRAM">ANAGRAM</option>
        <option value="READ_ALONG">READ_ALONG</option>
      </select>
    </div>
  );
};

export default Filter;
