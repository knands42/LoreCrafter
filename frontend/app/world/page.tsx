import React from 'react';
import Link from 'next/link';
import { FaPlus, FaGlobe } from 'react-icons/fa';

export const metadata = {
  title: 'Worlds - LoreCrafter',
  description: 'Build and manage your TTRPG worlds',
};

export default function WorldsPage() {
  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex justify-between items-center mb-8">
        <h1 className="text-3xl font-bold text-gray-900 dark:text-white">Your Worlds</h1>
        <Link 
          href="/world/create" 
          className="inline-flex items-center px-4 py-2 bg-purple-600 text-white rounded-md hover:bg-purple-700 transition-colors"
        >
          <FaPlus className="mr-2" />
          Create World
        </Link>
      </div>

      <div className="bg-gray-50 dark:bg-gray-800 rounded-lg p-8 text-center">
        <div className="flex justify-center mb-4">
          <FaGlobe className="h-16 w-16 text-purple-500" />
        </div>
        <h3 className="text-xl font-medium text-gray-900 dark:text-white mb-2">World Building Coming Soon</h3>
        <p className="text-gray-600 dark:text-gray-400 mb-6">
          Our world building tools are currently under development. Soon you'll be able to create rich, detailed worlds for your campaigns!
        </p>
        <div className="flex flex-col sm:flex-row justify-center gap-4">
          <Link 
            href="/campaign" 
            className="inline-flex items-center justify-center px-4 py-2 bg-purple-600 text-white rounded-md hover:bg-purple-700 transition-colors"
          >
            Explore Campaigns
          </Link>
        </div>
      </div>
    </div>
  );
}