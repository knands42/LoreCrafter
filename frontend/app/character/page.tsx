import React from 'react';
import Link from 'next/link';
import { FaPlus } from 'react-icons/fa';

export const metadata = {
  title: 'Characters - LoreCrafter',
  description: 'Manage your TTRPG characters',
};

export default function CharactersPage() {
  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex justify-between items-center mb-8">
        <h1 className="text-3xl font-bold text-gray-900 dark:text-white">Your Characters</h1>
        <Link 
          href="/character/create" 
          className="inline-flex items-center px-4 py-2 bg-purple-600 text-white rounded-md hover:bg-purple-700 transition-colors"
        >
          <FaPlus className="mr-2" />
          Create Character
        </Link>
      </div>

      <div className="bg-gray-50 dark:bg-gray-800 rounded-lg p-8 text-center">
        <h3 className="text-xl font-medium text-gray-900 dark:text-white mb-2">Character Management Coming Soon</h3>
        <p className="text-gray-600 dark:text-gray-400 mb-6">
          This feature is currently under development. Check back soon for updates!
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