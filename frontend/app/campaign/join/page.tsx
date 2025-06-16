'use client';

import React, { useState } from 'react';
import { useRouter } from 'next/navigation';
import Link from 'next/link';
import { joinCampaign } from '@/actions/campaignActions';
import { FaArrowRight, FaArrowLeft } from 'react-icons/fa';

export default function JoinCampaignPage() {
  const [inviteCode, setInviteCode] = useState('');
  const [error, setError] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError('');

    try {
      const member = await joinCampaign(inviteCode.trim());
      router.push(`/campaign/${member.campaign_id}`);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to join campaign');
      setIsLoading(false);
    }
  };

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="max-w-md mx-auto">
        <div className="mb-8">
          <Link 
            href="/campaign" 
            className="inline-flex items-center text-purple-600 hover:text-purple-700"
          >
            <FaArrowLeft className="mr-2" />
            Back to Campaigns
          </Link>
        </div>
        
        <div className="bg-white dark:bg-gray-800 rounded-lg shadow-md p-6">
          <h1 className="text-2xl font-bold text-gray-900 dark:text-white mb-6">Join a Campaign</h1>
          
          {error && (
            <div className="bg-red-50 dark:bg-red-900/30 p-4 rounded-md mb-6">
              <p className="text-red-800 dark:text-red-200">{error}</p>
            </div>
          )}
          
          <p className="text-gray-600 dark:text-gray-400 mb-6">
            Enter the invite code provided by the Game Master to join their campaign.
          </p>
          
          <form onSubmit={handleSubmit} className="space-y-6">
            <div>
              <label htmlFor="inviteCode" className="block text-sm font-medium text-gray-700 dark:text-gray-300">
                Invite Code
              </label>
              <input
                id="inviteCode"
                type="text"
                required
                value={inviteCode}
                onChange={(e) => setInviteCode(e.target.value)}
                className="mt-1 block w-full rounded-md border-gray-300 dark:border-gray-700 shadow-sm focus:border-purple-500 focus:ring-purple-500 dark:bg-gray-700 dark:text-white"
                placeholder="Enter invite code"
              />
            </div>
            
            <div>
              <button
                type="submit"
                disabled={isLoading}
                className="w-full inline-flex items-center justify-center px-4 py-2 bg-purple-600 text-white rounded-md hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-purple-500 disabled:bg-purple-400"
              >
                {isLoading ? 'Joining...' : 'Join Campaign'}
                {!isLoading && <FaArrowRight className="ml-2" />}
              </button>
            </div>
          </form>
          
          <div className="mt-8 pt-6 border-t border-gray-200 dark:border-gray-700">
            <p className="text-sm text-gray-500 dark:text-gray-400">
              Don't have an invite code? You can create your own campaign or ask a Game Master to invite you.
            </p>
            <div className="mt-4">
              <Link 
                href="/campaign/create" 
                className="text-purple-600 hover:text-purple-500 font-medium"
              >
                Create a new campaign
              </Link>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}