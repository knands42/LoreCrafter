'use client';

import React, {useEffect, useState} from 'react';
import {useRouter} from 'next/navigation';
import Link from 'next/link';
import {FaGlobe, FaPlus, FaSearch} from 'react-icons/fa';
import Button from '@/components/Button';
import Input from '@/components/Input';
import {worldApi} from '@/app/api/apiClient';
import {World} from '../api/types';

export default function WorldsList() {
  const router = useRouter();
  const [worlds, setWorlds] = useState<World[]>([]);
  const [searchQuery, setSearchQuery] = useState('');
  const [isSearching, setIsSearching] = useState(false);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  // Initial load - we'll just search with empty string to get some worlds
  useEffect(() => {
    const fetchInitialWorlds = async () => {
      try {
        const data = await worldApi.search('', 10);
        setWorlds(data);
      } catch (err) {
        console.error('Error fetching worlds:', err);
        setError('Failed to load worlds. Please try again.');
      } finally {
        setIsLoading(false);
      }
    };

    fetchInitialWorlds();
  }, []);

  const handleSearch = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsSearching(true);

    try {
      const data = await worldApi.search(searchQuery, 10);
      setWorlds(data);
      setError(null);
    } catch (err) {
      console.error('Error searching worlds:', err);
      setError('Failed to search worlds. Please try again.');
    } finally {
      setIsSearching(false);
    }
  };

  const handleWorldClick = (id: string) => {
    router.push(`/worlds/${id}`);
  };

  return (
    <div className="max-w-6xl mx-auto">
      <div className="flex flex-col md:flex-row md:items-center md:justify-between mb-8">
        <div>
          <h1 className="text-3xl font-bold text-gray-900 dark:text-white">Worlds</h1>
          <p className="mt-2 text-gray-600 dark:text-gray-400">
            View and manage your TTRPG worlds
          </p>
        </div>
        <div className="mt-4 md:mt-0">
          <Link href="/worlds/create">
            <Button>
              <FaPlus className="mr-2"/>
              Create World
            </Button>
          </Link>
        </div>
      </div>

      {/* Search Form */}
      <div className="mb-8">
        <form onSubmit={handleSearch} className="flex flex-col sm:flex-row gap-4">
          <div className="flex-grow">
            <Input
              label="Search Worlds"
              name="search"
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              placeholder="Search by name, history, or timeline..."
              className="mb-0"
            />
          </div>
          <div className="flex items-end">
            <Button type="submit" isLoading={isSearching}>
              <FaSearch className="mr-2"/>
              Search
            </Button>
          </div>
        </form>
      </div>

      {/* Error Message */}
      {error && (
        <div className="mb-6 bg-red-50 border-l-4 border-red-400 p-4">
          <div className="flex">
            <div>
              <p className="text-sm text-red-700">{error}</p>
            </div>
          </div>
        </div>
      )}

      {/* Loading State */}
      {isLoading && (
        <div className="flex justify-center items-center min-h-[40vh]">
          <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-indigo-500"></div>
        </div>
      )}

      {/* Worlds List */}
      {!isLoading && (
        <>
          {worlds.length === 0 ? (
            <div className="text-center py-12 bg-white dark:bg-gray-800 rounded-lg shadow">
              <FaGlobe className="mx-auto h-12 w-12 text-gray-400"/>
              <h3 className="mt-2 text-lg font-medium text-gray-900 dark:text-white">No worlds found</h3>
              <p className="mt-1 text-sm text-gray-500 dark:text-gray-400">
                {searchQuery
                  ? `No worlds match your search for "${searchQuery}"`
                  : "You haven't created any worlds yet"}
              </p>
              <div className="mt-6">
                <Link href="/worlds/create">
                  <Button>
                    <FaPlus className="mr-2"/>
                    Create your first world
                  </Button>
                </Link>
              </div>
            </div>
          ) : (
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
              {worlds.map((world) => (
                <div
                  key={world.id}
                  className="bg-white dark:bg-gray-800 overflow-hidden shadow rounded-lg cursor-pointer hover:shadow-md transition-shadow"
                  onClick={() => handleWorldClick(world.id)}
                >
                  <div className="px-4 py-5 sm:p-6">
                    <h3 className="text-lg font-medium text-gray-900 dark:text-white truncate">
                      {world.name}
                    </h3>
                    <div className="mt-1 text-sm text-gray-500 dark:text-gray-400">
                      {world.universe} • {world.world_theme} • {world.tone}
                    </div>
                    <div className="mt-3 text-sm text-gray-600 dark:text-gray-300 line-clamp-3">
                      {world.history}
                    </div>
                  </div>
                  <div className="bg-gray-50 dark:bg-gray-700 px-4 py-3 sm:px-6">
                    <div className="text-sm">
                        <span className="font-medium text-indigo-600 dark:text-indigo-400">
                          View timeline and details
                        </span>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          )}
        </>
      )}
    </div>
  );
}