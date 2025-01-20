import React from 'react';
import { Search, SortDesc } from 'lucide-react';

interface FiltersProps {
  search: string;
  onSearchChange: (value: string) => void;
  sortOrder: 'asc' | 'desc';
  onSortChange: (order: 'asc' | 'desc') => void;
}

export function Filters({
  search,
  onSearchChange,
  sortOrder,
  onSortChange,
}: FiltersProps) {
  return (
    <div className="flex flex-col sm:flex-row gap-4 mb-6">
      <div className="relative flex-1">
        <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 w-5 h-5" />
        <input
          type="text"
          value={search}
          onChange={(e) => onSearchChange(e.target.value)}
          placeholder="Search videos..."
          className="w-full pl-10 pr-4 py-2 rounded-lg border border-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
      </div>
      <button
        onClick={() => onSortChange(sortOrder === 'desc' ? 'asc' : 'desc')}
        className="flex items-center gap-2 px-4 py-2 rounded-lg bg-gray-100 hover:bg-gray-200"
      >
        <SortDesc className="w-5 h-5" />
        Sort by Date ({sortOrder === 'desc' ? 'Newest' : 'Oldest'})
      </button>
    </div>
  );
}