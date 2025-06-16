import React from 'react';
import Link from 'next/link';
import {Campaign} from '@/lib/types';
import {FaPlus} from 'react-icons/fa';

export const metadata = {
  title: 'Your Campaigns - LoreCrafter',
  description: 'View and manage your TTRPG campaigns',
};

export default async function CampaignsPage() {
  let campaigns: Campaign[] = [];
  let error = null;

  try {
    campaigns = await getCampaigns();
  } catch (err) {
    error = err instanceof Error ? err.message : 'Failed to load campaigns';
  }

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex justify-between items-center mb-8">
        <h1 className="text-3xl font-bold text-gray-900 dark:text-white">Your Campaigns</h1>
        <Link 
          href="/campaign/create" 
          className="inline-flex items-center px-4 py-2 bg-purple-600 text-white rounded-md hover:bg-purple-700 transition-colors"
        >
          <FaPlus className="mr-2" />
          Create Campaign
        </Link>
      </div>

      {error && (
        <div className="bg-red-50 dark:bg-red-900/30 p-4 rounded-md mb-6">
          <p className="text-red-800 dark:text-red-200">{error}</p>
        </div>
      )}

      {campaigns.length === 0 && !error ? (
        <div className="bg-gray-50 dark:bg-gray-800 rounded-lg p-8 text-center">
          <h3 className="text-xl font-medium text-gray-900 dark:text-white mb-2">No campaigns yet</h3>
          <p className="text-gray-600 dark:text-gray-400 mb-6">
            Create your first campaign or join an existing one with an invite code.
          </p>
          <div className="flex flex-col sm:flex-row justify-center gap-4">
            <Link 
              href="/campaign/create" 
              className="inline-flex items-center justify-center px-4 py-2 bg-purple-600 text-white rounded-md hover:bg-purple-700 transition-colors"
            >
              <FaPlus className="mr-2" />
              Create Campaign
            </Link>
            <Link 
              href="/campaign/join" 
              className="inline-flex items-center justify-center px-4 py-2 border border-purple-600 text-purple-600 rounded-md hover:bg-purple-50 dark:hover:bg-purple-900/20 transition-colors"
            >
              Join with Invite Code
            </Link>
          </div>
        </div>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {campaigns.map((campaign) => (
            <Link 
              key={campaign.id} 
              href={`/campaign/${campaign.id}`}
              className="block group"
            >
              <div className="bg-white dark:bg-gray-800 rounded-lg overflow-hidden shadow-md hover:shadow-lg transition-shadow">
                <div className="h-48 bg-gray-200 dark:bg-gray-700 relative">
                  {campaign.image_url ? (
                    <img 
                      src={campaign.image_url} 
                      alt={campaign.title} 
                      className="w-full h-full object-cover"
                    />
                  ) : (
                    <div className="w-full h-full flex items-center justify-center bg-gradient-to-br from-purple-500 to-indigo-600">
                      <span className="text-white text-2xl font-bold">
                        {campaign.title.charAt(0)}
                      </span>
                    </div>
                  )}
                  {!campaign.is_public && (
                    <div className="absolute top-2 right-2 bg-gray-800 text-white text-xs px-2 py-1 rounded">
                      Private
                    </div>
                  )}
                </div>
                <div className="p-4">
                  <h3 className="text-lg font-semibold text-gray-900 dark:text-white group-hover:text-purple-600 dark:group-hover:text-purple-400 transition-colors">
                    {campaign.title}
                  </h3>
                  {campaign.setting_summary && (
                    <p className="text-gray-600 dark:text-gray-400 mt-2 line-clamp-2">
                      {campaign.setting_summary}
                    </p>
                  )}
                  <div className="mt-4 text-sm text-gray-500 dark:text-gray-500">
                    Created {new Date(campaign.created_at).toLocaleDateString()}
                  </div>
                </div>
              </div>
            </Link>
          ))}
        </div>
      )}
    </div>
  );
}