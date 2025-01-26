import React from 'react';

function Pagination({ 
  currentPage, 
  totalPages, 
  onPreviousPage, 
  onNextPage 
}) {
  return (
    <div className="pagination">
      <button 
        onClick={onPreviousPage} 
        disabled={currentPage === 1}
        className="pagination-button"
      >
        ← Previous
      </button>
      <span className="pagination-text">
        Page {currentPage} of {totalPages}
      </span>
      <button 
        onClick={onNextPage} 
        disabled={currentPage === totalPages}
        className="pagination-button"
      >
        Next →
      </button>
    </div>
  );
}

export default Pagination;