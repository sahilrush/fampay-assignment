import { TGetVideo } from './types';

export const mockVideos: TGetVideo[] = [
  {
    id: 1,
    title: 'Building a React Application',
    description: 'Learn how to build a modern React application from scratch with TypeScript and Tailwind CSS.',
    publishedAt: '2024-03-10T10:00:00Z',
    thumbnails: 'https://images.unsplash.com/photo-1633356122544-f134324a6cee?w=320&h=180&fit=crop'
  },
  {
    id: 2,
    title: 'TypeScript Tips and Tricks',
    description: 'Advanced TypeScript features and patterns for better code quality.',
    publishedAt: '2024-03-09T15:30:00Z',
    thumbnails: 'https://images.unsplash.com/photo-1587620962725-abab7fe55159?w=320&h=180&fit=crop'
  }
];