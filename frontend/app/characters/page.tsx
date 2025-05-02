'use client';

import React, {useEffect, useState} from 'react';
import {useRouter} from 'next/navigation';
import Link from 'next/link';
import {FaPlus, FaSearch, FaUser} from 'react-icons/fa';
import Button from '@/components/Button';
import Input from '@/components/Input';
import {characterApi} from '@/app/api/apiClient';
import {Character} from '../api/types';

export default function CharactersList() {
  const router = useRouter();
  const [characters, setCharacters] = useState<Character[]>([]);
  const [searchQuery, setSearchQuery] = useState('');
  const [isSearching, setIsSearching] = useState(false);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  // Initial load - we'll just search with empty string to get some characters
  useEffect(() => {
    const fetchInitialCharacters = async () => {
      try {
        const data = await characterApi.search('', 10);
        setCharacters(data);
      } catch (err) {
        console.error('Error fetching characters:', err);
        setError('Failed to load characters. Please try again.');
      } finally {
        setIsLoading(false);
      }
    };

    fetchInitialCharacters();
  }, []);

  const handleSearch = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsSearching(true);

    try {
      const data = await characterApi.search(searchQuery, 10);
      setCharacters(data);
      setError(null);
    } catch (err) {
      console.error('Error searching characters:', err);
      setError('Failed to search characters. Please try again.');
    } finally {
      setIsSearching(false);
    }
  };

  const handleCharacterClick = (id: string) => {
    router.push(`/characters/${id}`);
  };

  return (
    <div className="max-w-6xl mx-auto">
      <div className="flex flex-col md:flex-row md:items-center md:justify-between mb-8">
        <div>
          <h1 className="text-3xl font-bold text-gray-900 dark:text-white">Characters</h1>
          <p className="mt-2 text-gray-600 dark:text-gray-400">
            View and manage your TTRPG characters
          </p>
        </div>
        <div className="mt-4 md:mt-0">
          <Link href="/characters/create">
            <Button>
              <FaPlus className="mr-2"/>
              Create Character
            </Button>
          </Link>
        </div>
      </div>

      {/* Search Form */}
      <div className="mb-8">
        <form onSubmit={handleSearch} className="flex flex-col sm:flex-row gap-4">
          <div className="flex-grow">
            <Input
              label="Search Characters"
              name="search"
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              placeholder="Search by name, race, or backstory..."
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
          <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-purple-500"></div>
        </div>
      )}

      {/* Characters List */}
      {!isLoading && (
        <>
          {characters.length === 0 ? (
            <div className="text-center py-12 bg-white dark:bg-gray-800 rounded-lg shadow">
              <FaUser className="mx-auto h-12 w-12 text-gray-400"/>
              <h3 className="mt-2 text-lg font-medium text-gray-900 dark:text-white">No characters found</h3>
              <p className="mt-1 text-sm text-gray-500 dark:text-gray-400">
                {searchQuery
                  ? `No characters match your search for "${searchQuery}"`
                  : "You haven't created any characters yet"}
              </p>
              <div className="mt-6">
                <Link href="/characters/create">
                  <Button>
                    <FaPlus className="mr-2"/>
                    Create your first character
                  </Button>
                </Link>
              </div>
            </div>
          ) : (
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {characters.map((character) => (
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
                    <div className="text-sm text-gray-500 dark:text-gray-400">
                        <span className="font-medium text-purple-600 dark:text-purple-400">
                          {character.world_theme}
                        </span>
                      {' • '}
                      <span className="font-medium text-indigo-600 dark:text-indigo-400">
                          {character.tone}
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