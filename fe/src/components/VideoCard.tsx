import React from 'react';
import { Calendar } from 'lucide-react';
import { TGetVideo } from '../types';

interface VideoCardProps {
  video: TGetVideo;
}

export function VideoCard({ video }: VideoCardProps) {
  return (
    <div className="bg-white rounded-lg shadow-md overflow-hidden">
      <img
        src={video.thumbnails}
        alt={video.title}
        className="w-full h-48 object-cover"
      />
      <div className="p-4">
        <h3 className="font-semibold text-lg mb-2 line-clamp-2">{video.title}</h3>
        <p className="text-gray-600 text-sm mb-3 line-clamp-3">{video.description}</p>
        <div className="flex items-center text-gray-500 text-sm">
          <Calendar className="w-4 h-4 mr-1" />
          {new Date(video.publishedAt).toLocaleDateString()}
        </div>
      </div>
    </div>
  );
}