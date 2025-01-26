import React, { useState } from 'react';
import SearchBar from './components/SearchBar';
import ResultsList from './components/ResultsList';
import Pagination from './components/Pagination';
import './index.css';

function App() {
  const [query, setQuery] = useState('');
  const [results, setResults] = useState([]);
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
      // No filter logic, just use the query directly
      const response = await fetch(`/search?query=${query}`);
      const data = await response.json();
      setResults(data);
      setQuery(''); // Clear the search bar after generating results
    } catch (error) {
      console.error("Error fetching search results:", error);
      setResults([]);
    } finally {
      setLoading(false);
    }
  };

  const indexOfLastItem = currentPage * itemsPerPage;
  const indexOfFirstItem = indexOfLastItem - itemsPerPage;
  const currentResults = results.slice(indexOfFirstItem, indexOfLastItem);
  const totalPages = Math.ceil(results.length / itemsPerPage);

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
          <ResultsList 
            results={currentResults} 
            loading={loading} 
            searched={searched} 
          />

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
    </div>
  );
}

export default App;
