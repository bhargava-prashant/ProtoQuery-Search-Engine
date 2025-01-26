import React from 'react';

function SearchBar({ query, setQuery, onSearch, loading }) {
  return (
    <div className="search-bar-container">
      <h1 className="search-title">
        Question Search WebApp
      </h1>
      <form onSubmit={onSearch} className="search-form">
        <div className="search-input-wrapper">
          <input 
            type="text" 
            placeholder={loading ? 'Searching...' : 'Enter your search query'}
            value={query} 
            onChange={(e) => setQuery(e.target.value)}
            className="search-input"
          />
        </div>
        <button 
          type="submit" 
          disabled={loading}
          className="search-button"
        >
          {loading ? 'Searching...' : 'Search'}
        </button>
      </form>
    </div>
  );
}

export default SearchBar;