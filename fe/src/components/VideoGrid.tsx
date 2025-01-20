import React from 'react';
import { TGetVideo } from '../types';
import { VideoCard } from './VideoCard';

interface VideoGridProps {
  videos: TGetVideo[];
}

export function VideoGrid({ videos }: VideoGridProps) {
  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
      {videos.map((video) => (
        <VideoCard key={video.id} video={video} />
      ))}
    </div>
  );
}