'use client';

import React, { useState } from 'react';
import { useRouter } from 'next/navigation';
import { createCampaign } from '@/actions/campaignActions';
import { CampaignCreationRequest } from '@/lib/types';
import { FaSave, FaTimes } from 'react-icons/fa';

export default function CreateCampaignPage() {
  const [title, setTitle] = useState('');
  const [settingSummary, setSettingSummary] = useState('');
  const [setting, setSetting] = useState('');
  const [imageUrl, setImageUrl] = useState('');
  const [isPublic, setIsPublic] = useState(false);
  const [error, setError] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError('');

    const campaignData: CampaignCreationRequest = {
      title,
      setting_summary: settingSummary,
      setting,
      image_url: imageUrl,
      is_public: isPublic,
    };

    try {
      const campaign = await createCampaign(campaignData);
      router.push(`/campaign/${campaign.id}`);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to create campaign');
      setIsLoading(false);
    }
  };

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="max-w-3xl mx-auto">
        <div className="flex justify-between items-center mb-8">
          <h1 className="text-3xl font-bold text-gray-900 dark:text-white">Create New Campaign</h1>
          <button
            onClick={() => router.back()}
            className="inline-flex items-center px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-md text-gray-700 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-800"
          >
            <FaTimes className="mr-2" />
            Cancel
          </button>
        </div>

        {error && (
          <div className="bg-red-50 dark:bg-red-900/30 p-4 rounded-md mb-6">
            <p className="text-red-800 dark:text-red-200">{error}</p>
          </div>
        )}

        <form onSubmit={handleSubmit} className="bg-white dark:bg-gray-800 rounded-lg shadow-md p-6">
          <div className="space-y-6">
            <div>
              <label htmlFor="title" className="block text-sm font-medium text-gray-700 dark:text-gray-300">
                Campaign Title <span className="text-red-500">*</span>
              </label>
              <input
                id="title"
                type="text"
                required
                value={title}
                onChange={(e) => setTitle(e.target.value)}
                className="mt-1 block w-full rounded-md border-gray-300 dark:border-gray-700 shadow-sm focus:border-purple-500 focus:ring-purple-500 dark:bg-gray-700 dark:text-white"
                placeholder="Enter a name for your campaign"
              />
            </div>

            <div>
              <label htmlFor="settingSummary" className="block text-sm font-medium text-gray-700 dark:text-gray-300">
                Setting Summary
              </label>
              <input
                id="settingSummary"
                type="text"
                value={settingSummary}
                onChange={(e) => setSettingSummary(e.target.value)}
                className="mt-1 block w-full rounded-md border-gray-300 dark:border-gray-700 shadow-sm focus:border-purple-500 focus:ring-purple-500 dark:bg-gray-700 dark:text-white"
                placeholder="A brief description of your campaign setting"
              />
              <p className="mt-1 text-sm text-gray-500 dark:text-gray-400">
                This will be displayed on the campaign card.
              </p>
            </div>

            <div>
              <label htmlFor="setting" className="block text-sm font-medium text-gray-700 dark:text-gray-300">
                Campaign Setting
              </label>
              <textarea
                id="setting"
                rows={6}
                value={setting}
                onChange={(e) => setSetting(e.target.value)}
                className="mt-1 block w-full rounded-md border-gray-300 dark:border-gray-700 shadow-sm focus:border-purple-500 focus:ring-purple-500 dark:bg-gray-700 dark:text-white"
                placeholder="Describe your campaign world, lore, and setting in detail"
              />
            </div>

            <div>
              <label htmlFor="imageUrl" className="block text-sm font-medium text-gray-700 dark:text-gray-300">
                Cover Image URL
              </label>
              <input
                id="imageUrl"
                type="url"
                value={imageUrl}
                onChange={(e) => setImageUrl(e.target.value)}
                className="mt-1 block w-full rounded-md border-gray-300 dark:border-gray-700 shadow-sm focus:border-purple-500 focus:ring-purple-500 dark:bg-gray-700 dark:text-white"
                placeholder="https://example.com/image.jpg"
              />
              <p className="mt-1 text-sm text-gray-500 dark:text-gray-400">
                Provide a URL to an image that represents your campaign.
              </p>
            </div>

            <div className="flex items-center">
              <input
                id="isPublic"
                type="checkbox"
                checked={isPublic}
                onChange={(e) => setIsPublic(e.target.checked)}
                className="h-4 w-4 text-purple-600 focus:ring-purple-500 border-gray-300 rounded"
              />
              <label htmlFor="isPublic" className="ml-2 block text-sm text-gray-700 dark:text-gray-300">
                Make this campaign public
              </label>
            </div>
            <p className="text-sm text-gray-500 dark:text-gray-400">
              Public campaigns can be discovered by other users. Private campaigns are only accessible via invite.
            </p>
          </div>

          <div className="mt-8 flex justify-end">
            <button
              type="submit"
              disabled={isLoading}
              className="inline-flex items-center px-4 py-2 bg-purple-600 text-white rounded-md hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-purple-500 disabled:bg-purple-400"
            >
              <FaSave className="mr-2" />
              {isLoading ? 'Creating...' : 'Create Campaign'}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}