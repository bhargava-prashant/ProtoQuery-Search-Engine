import React, { useState } from 'react';
import SearchBar from './components/SearchBar';
import ResultsList from './components/ResultsList';
import Pagination from './components/Pagination';
import './index.css';

function App() {
  const [query, setQuery] = useState('');
  const [results, setResults] = useState([]); // Initialize as empty array
  const [searched, setSearched] = useState(false);
  const [loading, setLoading] = useState(false);
  const [currentPage, setCurrentPage] = useState(1);
  const itemsPerPage = 25;

  const handleSearch = async (e) => {
    e.preventDefault();
    setSearched(true);
    setLoading(true);
    setCurrentPage(1);

    try {
      const response = await fetch(`/search?query=${query}`);
      const data = await response.json();
      
      // If no results found, set results to empty array
      if (!data || data.length === 0) {
        setResults([]);
      } else {
        setResults(data);
      }
      
      setQuery(''); // Clear the search bar after generating results
    } catch (error) {
      console.error("Error fetching search results:", error);
      setResults([]); // Ensure results is an empty array in case of error
    } finally {
      setLoading(false);
    }
  };

  // Prevent slice error if results are empty
  const indexOfLastItem = currentPage * itemsPerPage;
  const indexOfFirstItem = indexOfLastItem - itemsPerPage;
  const currentResults = results.slice(indexOfFirstItem, indexOfLastItem);
  const totalPages = Math.ceil(results.length / itemsPerPage);

  const handleBackToList = () => {
    setSearched(false); // Optional: Reset searched state to show the initial message
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-100 to-gray-200 flex items-center justify-center p-6">
      <div className="w-full max-w-4xl bg-white shadow-2xl rounded-2xl overflow-hidden">
        <SearchBar 
          query={query} 
          setQuery={setQuery} 
          onSearch={handleSearch} 
          loading={loading} 
        />

        <div className="p-6">
          {!searched ? (
            <div className="initial-message"></div>
          ) : (
            <>
              {results.length === 0 ? (
                <div className="no-results-message">No results found. Try a different query.</div>
              ) : (
                <ResultsList 
                  results={currentResults} 
                  loading={loading} 
                  searched={searched} 
                  onBackToList={handleBackToList} // Pass back function
                />
              )}
            </>
          )}

          {results.length > itemsPerPage && (
            <Pagination 
              currentPage={currentPage}
              totalPages={totalPages}
              onPreviousPage={() => setCurrentPage(currentPage - 1)}
              onNextPage={() => setCurrentPage(currentPage + 1)}
            />
          )}
        </div>
      </div>

      {/* Footer with copyright */}
      <footer className="absolute bottom-0 w-full text-center py-4 bg-gray-800 text-white">
        <p>&copy; 2025 Prashant Bhargava. All rights reserved.</p>
      </footer>
    </div>
  );
}

export default App;