'use client';

import React, {useState} from 'react';
import {useRouter} from 'next/navigation';
import {FaGlobe, FaSearch, FaUser} from 'react-icons/fa';
import Button from '@/components/Button';
import Input from '@/components/Input';
import Select from '@/components/Select';
import {characterApi, worldApi} from '@/app/api/apiClient';
import {Character, World} from '../api/types';

type SearchType = 'all' | 'characters' | 'worlds';

const searchTypeOptions = [
  {value: 'all', label: 'All'},
  {value: 'characters', label: 'Characters'},
  {value: 'worlds', label: 'Worlds'},
];

export default function SearchPage() {
  const router = useRouter();
  const [searchQuery, setSearchQuery] = useState('');
  const [searchType, setSearchType] = useState<SearchType>('all');
  const [isSearching, setIsSearching] = useState(false);
  const [characterResults, setCharacterResults] = useState<Character[]>([]);
  const [worldResults, setWorldResults] = useState<World[]>([]);
  const [hasSearched, setHasSearched] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleSearch = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!searchQuery.trim()) return;

    setIsSearching(true);
    setError(null);
    setHasSearched(true);

    try {
      // Search characters if type is 'all' or 'characters'
      if (searchType === 'all' || searchType === 'characters') {
        const characters = await characterApi.search(searchQuery, 10);
        setCharacterResults(characters);
      } else {
        setCharacterResults([]);
      }

      // Search worlds if type is 'all' or 'worlds'
      if (searchType === 'all' || searchType === 'worlds') {
        const worlds = await worldApi.search(searchQuery, 10);
        setWorldResults(worlds);
      } else {
        setWorldResults([]);
      }
    } catch (err) {
      console.error('Error searching:', err);
      setError('Failed to search. Please try again.');
    } finally {
      setIsSearching(false);
    }
  };

  const handleCharacterClick = (id: string) => {
    router.push(`/characters/${id}`);
  };

  const handleWorldClick = (id: string) => {
    router.push(`/worlds/${id}`);
  };

  const handleSearchTypeChange = (value: string) => {
    setSearchType(value as SearchType);
  };

  return (
    <div className="max-w-6xl mx-auto">
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-gray-900 dark:text-white">Search</h1>
        <p className="mt-2 text-gray-600 dark:text-gray-400">
          Search for characters and worlds across your TTRPG content
        </p>
      </div>

      {/* Search Form */}
      <div className="bg-white dark:bg-gray-800 shadow rounded-lg p-6 mb-8">
        <form onSubmit={handleSearch} className="space-y-4">
          <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
            <div className="md:col-span-3">
              <Input
                label="Search Query"
                name="search"
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
                placeholder="Enter search terms..."
                required
              />
            </div>
            <div>
              <Select
                label="Search Type"
                name="searchType"
                options={searchTypeOptions}
                value={searchType}
                onChange={handleSearchTypeChange}
              />
            </div>
          </div>
          <div className="flex justify-end">
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

      {/* Search Results */}
      {hasSearched && !isSearching && (
        <div className="space-y-8">
          {/* No Results Message */}
          {characterResults.length === 0 && worldResults.length === 0 && (
            <div className="text-center py-12 bg-white dark:bg-gray-800 rounded-lg shadow">
              <FaSearch className="mx-auto h-12 w-12 text-gray-400"/>
              <h3 className="mt-2 text-lg font-medium text-gray-900 dark:text-white">No results found</h3>
              <p className="mt-1 text-sm text-gray-500 dark:text-gray-400">
                No matches found for "{searchQuery}". Try a different search term or type.
              </p>
            </div>
          )}

          {/* Character Results */}
          {characterResults.length > 0 && (
            <div>
              <h2 className="text-xl font-bold text-gray-900 dark:text-white mb-4 flex items-center">
                <FaUser className="mr-2"/>
                Characters ({characterResults.length})
              </h2>
              <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                {characterResults.map((character) => (
                  <div
                    key={character.id}
                    className="bg-white dark:bg-gray-800 overflow-hidden shadow rounded-lg cursor-pointer hover:shadow-md transition-shadow"
                    onClick={() => handleCharacterClick(character.id)}
                  >
                    <div className="px-4 py-5 sm:p-6">
                      <h3 className="text-lg font-medium text-gray-900 dark:text-white truncate">
                        {character.name}
                      </h3>
                      <div className="mt-1 text-sm text-gray-500 dark:text-gray-400">
                        {character.race} • {character.gender} • {character.universe}
                      </div>
                      <div className="mt-3 text-sm text-gray-600 dark:text-gray-300 line-clamp-3">
                        {character.backstory}
                      </div>
                    </div>
                    <div className="bg-gray-50 dark:bg-gray-700 px-4 py-3 sm:px-6">
                      <div className="text-sm">
                          <span className="font-medium text-purple-600 dark:text-purple-400">
                            View character details
                          </span>
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            </div>
          )}

          {/* World Results */}
          {worldResults.length > 0 && (
            <div>
              <h2 className="text-xl font-bold text-gray-900 dark:text-white mb-4 flex items-center">
                <FaGlobe className="mr-2"/>
                Worlds ({worldResults.length})
              </h2>
              <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                {worldResults.map((world) => (
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
                            View world details
                          </span>
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            </div>
          )}
        </div>
      )}
    </div>
  );
}