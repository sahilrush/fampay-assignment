import React, { useState, useEffect } from 'react';
import { Youtube } from 'lucide-react';
import { VideoGrid } from './components/VideoGrid';
import { Pagination } from './components/Pagination';
import { Filters } from './components/Filters';
import { TGetVideo, TGetVideosResponse } from './types';
import { mockVideos } from './mockData';

// Set this to false when your backend is ready
const USE_MOCK_DATA = false;

function App() {
  const [videos, setVideos] = useState<TGetVideo[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [page, setPage] = useState(1);
  const [hasNextPage, setHasNextPage] = useState(false);
  const [search, setSearch] = useState('');
  const [sortOrder, setSortOrder] = useState<'asc' | 'desc'>('desc');

  useEffect(() => {
    const fetchVideos = async () => {
      try {
        setLoading(true);
        setError(null);

        if (USE_MOCK_DATA) {
          // Simulate API delay
          await new Promise(resolve => setTimeout(resolve, 500));

          // Filter mock data based on search
          let filteredVideos = mockVideos.filter(video =>
            video.title.toLowerCase().includes(search.toLowerCase()) ||
            video.description.toLowerCase().includes(search.toLowerCase())
          );

          // Sort videos
          filteredVideos.sort((a, b) => {
            const dateA = new Date(a.publishedAt).getTime();
            const dateB = new Date(b.publishedAt).getTime();
            return sortOrder === 'desc' ? dateB - dateA : dateA - dateB;
          });

          setVideos(filteredVideos);
          setHasNextPage(false);
          return;
        }

        const params = new URLSearchParams({
          page: page.toString(),
          search: search,
          sort: sortOrder,
          limit: "6"
        });
        
        const response = await fetch(`http://localhost:8080/videos?${params}`);
        if (!response.ok) {
          const errorText = await response.text();
          throw new Error(
            `Failed to fetch videos: ${response.status} ${response.statusText}\n${errorText}`
          );
        }
        
        const data: TGetVideosResponse = await response.json();
        setVideos(data.videos);
        setHasNextPage(data.videos.length === data.limit);
      } catch (err) {
        console.error('Error fetching videos:', err);
        setError(err instanceof Error ? err.message : 'An error occurred while fetching videos');
      } finally {
        setLoading(false);
      }
    };

    fetchVideos();
  }, [page, search, sortOrder]);

  if (error) {
    return (
      <div className="min-h-screen bg-gray-50 flex items-center justify-center">
        <div className="text-center max-w-lg mx-auto p-6">
          <p className="text-red-500 mb-4 whitespace-pre-wrap">{error}</p>
          <button
            onClick={() => window.location.reload()}
            className="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition-colors"
          >
            Retry
          </button>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-50">
      <header className="bg-white shadow-sm">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
          <div className="flex items-center gap-2">
            <Youtube className="w-8 h-8 text-red-500" />
            <h1 className="text-xl font-semibold">YouTube Video Dashboard</h1>
          </div>
        </div>
      </header>

      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <Filters
          search={search}
          onSearchChange={setSearch}
          sortOrder={sortOrder}
          onSortChange={setSortOrder}
        />

        {loading ? (
          <div className="flex items-center justify-center min-h-[400px]">
            <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-gray-900"></div>
          </div>
        ) : (
          <>
            {videos.length > 0 ? (
              <VideoGrid videos={videos} />
            ) : (
              <div className="text-center py-12 text-gray-500">
                No videos found
              </div>
            )}
            {!USE_MOCK_DATA && (
              <Pagination
                currentPage={page}
                hasNextPage={hasNextPage}
                hasPrevPage={page > 1}
                onPageChange={setPage}
              />
            )}
          </>
        )}
      </main>
    </div>
  );
}

export default App;