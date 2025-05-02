'use client';

import React, {useEffect, useState} from 'react';
import {useRouter} from 'next/navigation';
import Link from 'next/link';
import {FaDungeon, FaGlobe, FaPlus, FaSearch, FaUser} from 'react-icons/fa';
import Button from '@/components/Button';
import Input from '@/components/Input';
import {campaignApi} from '@/app/api/apiClient';
import {Campaign} from '../api/types';

export default function CampaignsList() {
  const router = useRouter();
  const [campaigns, setCampaigns] = useState<Campaign[]>([]);
  const [searchQuery, setSearchQuery] = useState('');
  const [isSearching, setIsSearching] = useState(false);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  // Initial load - we'll just search with empty string to get some campaigns
  useEffect(() => {
    const fetchInitialCampaigns = async () => {
      try {
        // Note: The API doesn't have a search endpoint for campaigns, so we'll just get all campaigns
        // In a real application, you would implement proper search functionality
        const data = await campaignApi.getAll();
        setCampaigns(data);
      } catch (err) {
        console.error('Error fetching campaigns:', err);
        setError('Failed to load campaigns. Please try again.');
      } finally {
        setIsLoading(false);
      }
    };

    fetchInitialCampaigns();
  }, []);

  const handleCampaignClick = (id: string) => {
    router.push(`/campaigns/${id}`);
  };

  // Note: This is a mock search function since the API doesn't support campaign search
  // In a real application, you would call the API with the search query
  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault();
    if (!searchQuery.trim()) return;

    setIsSearching(true);

    // Mock search by filtering the campaigns client-side
    const filteredCampaigns = campaigns.filter(campaign =>
      campaign.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
      campaign.campaign.toLowerCase().includes(searchQuery.toLowerCase())
    );

    setCampaigns(filteredCampaigns);
    setIsSearching(false);
  };

  return (
    <div className="max-w-6xl mx-auto">
      <div className="flex flex-col md:flex-row md:items-center md:justify-between mb-8">
        <div>
          <h1 className="text-3xl font-bold text-gray-900 dark:text-white">Campaigns</h1>
          <p className="mt-2 text-gray-600 dark:text-gray-400">
            View and manage your TTRPG campaigns
          </p>
        </div>
        <div className="mt-4 md:mt-0">
          <Link href="/campaigns/create">
            <Button>
              <FaPlus className="mr-2"/>
              Create Campaign
            </Button>
          </Link>
        </div>
      </div>

      {/* Search Form */}
      <div className="mb-8">
        <form onSubmit={handleSearch} className="flex flex-col sm:flex-row gap-4">
          <div className="flex-grow">
            <Input
              label="Search Campaigns"
              name="search"
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              placeholder="Search by name or description..."
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
          <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-blue-500"></div>
        </div>
      )}

      {/* Campaigns List */}
      {!isLoading && (
        <>
          {campaigns.length === 0 ? (
            <div className="text-center py-12 bg-white dark:bg-gray-800 rounded-lg shadow">
              <FaDungeon className="mx-auto h-12 w-12 text-gray-400"/>
              <h3 className="mt-2 text-lg font-medium text-gray-900 dark:text-white">No campaigns found</h3>
              <p className="mt-1 text-sm text-gray-500 dark:text-gray-400">
                {searchQuery
                  ? `No campaigns match your search for "${searchQuery}"`
                  : "You haven't created any campaigns yet"}
              </p>
              <div className="mt-6">
                <Link href="/campaigns/create">
                  <Button>
                    <FaPlus className="mr-2"/>
                    Create your first campaign
                  </Button>
                </Link>
              </div>
            </div>
          ) : (
            <div className="grid grid-cols-1 gap-6">
              {campaigns.map((campaign) => (
                <div
                  key={campaign.id}
                  className="bg-white dark:bg-gray-800 overflow-hidden shadow rounded-lg cursor-pointer hover:shadow-md transition-shadow"
                  onClick={() => handleCampaignClick(campaign.id)}
                >
                  <div className="px-4 py-5 sm:p-6">
                    <h3 className="text-lg font-medium text-gray-900 dark:text-white truncate">
                      {campaign.name}
                    </h3>
                    <div className="mt-1 text-sm text-gray-500 dark:text-gray-400">
                      {campaign.universe} • {campaign.world_theme} • {campaign.tone}
                    </div>
                    <div className="mt-3 text-sm text-gray-600 dark:text-gray-300 line-clamp-3">
                      {campaign.campaign}
                    </div>

                    {/* Linked World and Characters */}
                    <div className="mt-4 flex flex-wrap gap-2">
                      {campaign.linked_world_id && (
                        <span
                          className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-indigo-100 text-indigo-800 dark:bg-indigo-800 dark:text-indigo-100">
                            <FaGlobe className="mr-1"/>
                            Linked World
                          </span>
                      )}

                      {campaign.linked_character_ids && campaign.linked_character_ids.length > 0 && (
                        <span
                          className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-purple-100 text-purple-800 dark:bg-purple-800 dark:text-purple-100">
                            <FaUser className="mr-1"/>
                          {campaign.linked_character_ids.length} Character{campaign.linked_character_ids.length !== 1 ? 's' : ''}
                          </span>
                      )}
                    </div>
                  </div>
                  <div className="bg-gray-50 dark:bg-gray-700 px-4 py-3 sm:px-6">
                    <div className="text-sm">
                        <span className="font-medium text-blue-600 dark:text-blue-400">
                          View campaign details and hidden elements
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

// Note: In a real application, you would add this method to the apiClient.ts file
// campaignApi.getAll = async () => {
//   const response = await apiClient.get('/campaigns');
//   return response.data;
// };
