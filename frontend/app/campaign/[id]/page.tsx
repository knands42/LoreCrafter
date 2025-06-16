import React from 'react';
import Link from 'next/link';
import { notFound } from 'next/navigation';
import { getCampaign, getCampaignMembers } from '@/actions/campaignActions';
import { FaUsers, FaLink, FaArrowLeft } from 'react-icons/fa';

interface CampaignDetailPageProps {
  params: {
    id: string;
  };
}

export default async function CampaignDetailPage({ params }: CampaignDetailPageProps) {
  const { id } = params;
  
  try {
    const [campaign, members] = await Promise.all([
      getCampaign(id),
      getCampaignMembers(id),
    ]);
    
    const gamemaster = members.find(member => member.role === 'gm');
    const players = members.filter(member => member.role === 'player');
    
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="mb-8">
          <Link 
            href="/campaign" 
            className="inline-flex items-center text-purple-600 hover:text-purple-700"
          >
            <FaArrowLeft className="mr-2" />
            Back to Campaigns
          </Link>
        </div>
        
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
          {/* Main Content */}
          <div className="lg:col-span-2 space-y-6">
            <div className="bg-white dark:bg-gray-800 rounded-lg shadow-md overflow-hidden">
              {campaign.image_url && (
                <div className="h-64 bg-gray-200 dark:bg-gray-700">
                  <img 
                    src={campaign.image_url} 
                    alt={campaign.title} 
                    className="w-full h-full object-cover"
                  />
                </div>
              )}
              
              <div className="p-6">
                <div className="flex justify-between items-start">
                  <h1 className="text-3xl font-bold text-gray-900 dark:text-white">{campaign.title}</h1>
                  {!campaign.is_public && (
                    <span className="bg-gray-100 dark:bg-gray-700 text-gray-800 dark:text-gray-300 text-xs px-2 py-1 rounded">
                      Private
                    </span>
                  )}
                </div>
                
                {campaign.setting_summary && (
                  <p className="mt-2 text-gray-600 dark:text-gray-400 text-lg">
                    {campaign.setting_summary}
                  </p>
                )}
                
                <div className="mt-6">
                  <h2 className="text-xl font-semibold text-gray-900 dark:text-white mb-3">Campaign Setting</h2>
                  <div className="prose dark:prose-invert max-w-none">
                    {campaign.setting ? (
                      <p className="text-gray-600 dark:text-gray-400 whitespace-pre-line">
                        {campaign.setting}
                      </p>
                    ) : (
                      <p className="text-gray-500 dark:text-gray-500 italic">
                        No campaign setting has been provided yet.
                      </p>
                    )}
                  </div>
                </div>
              </div>
            </div>
            
            {/* Additional sections like timeline, characters, etc. would go here */}
          </div>
          
          {/* Sidebar */}
          <div className="space-y-6">
            {/* Campaign Info */}
            <div className="bg-white dark:bg-gray-800 rounded-lg shadow-md p-6">
              <h2 className="text-xl font-semibold text-gray-900 dark:text-white mb-4">Campaign Info</h2>
              
              <div className="space-y-4">
                <div>
                  <h3 className="text-sm font-medium text-gray-500 dark:text-gray-400">Created</h3>
                  <p className="mt-1 text-gray-900 dark:text-white">
                    {new Date(campaign.created_at).toLocaleDateString()}
                  </p>
                </div>
                
                <div>
                  <h3 className="text-sm font-medium text-gray-500 dark:text-gray-400">Last Updated</h3>
                  <p className="mt-1 text-gray-900 dark:text-white">
                    {new Date(campaign.updated_at).toLocaleDateString()}
                  </p>
                </div>
                
                {campaign.invite_code && (
                  <div>
                    <h3 className="text-sm font-medium text-gray-500 dark:text-gray-400">Invite Code</h3>
                    <div className="mt-1 flex items-center">
                      <code className="bg-gray-100 dark:bg-gray-700 px-2 py-1 rounded text-sm font-mono text-gray-900 dark:text-white">
                        {campaign.invite_code}
                      </code>
                      <button 
                        className="ml-2 text-purple-600 hover:text-purple-700"
                        onClick={() => {
                          navigator.clipboard.writeText(campaign.invite_code || '');
                        }}
                      >
                        <FaLink className="h-4 w-4" />
                      </button>
                    </div>
                    <p className="mt-1 text-xs text-gray-500 dark:text-gray-400">
                      Share this code with players you want to invite
                    </p>
                  </div>
                )}
              </div>
            </div>
            
            {/* Members */}
            <div className="bg-white dark:bg-gray-800 rounded-lg shadow-md p-6">
              <div className="flex items-center justify-between mb-4">
                <h2 className="text-xl font-semibold text-gray-900 dark:text-white">Members</h2>
                <FaUsers className="h-5 w-5 text-gray-400" />
              </div>
              
              {gamemaster && gamemaster.user && (
                <div className="mb-4">
                  <h3 className="text-sm font-medium text-gray-500 dark:text-gray-400 mb-2">Game Master</h3>
                  <div className="flex items-center">
                    <div className="flex-shrink-0">
                      {gamemaster.user.avatar_url ? (
                        <img 
                          src={gamemaster.user.avatar_url} 
                          alt={gamemaster.user.username} 
                          className="h-8 w-8 rounded-full"
                        />
                      ) : (
                        <div className="h-8 w-8 rounded-full bg-purple-600 flex items-center justify-center">
                          <span className="text-white text-sm font-medium">
                            {gamemaster.user.username.charAt(0).toUpperCase()}
                          </span>
                        </div>
                      )}
                    </div>
                    <div className="ml-3">
                      <p className="text-sm font-medium text-gray-900 dark:text-white">
                        {gamemaster.user.username}
                      </p>
                    </div>
                  </div>
                </div>
              )}
              
              {players.length > 0 && (
                <div>
                  <h3 className="text-sm font-medium text-gray-500 dark:text-gray-400 mb-2">Players ({players.length})</h3>
                  <ul className="space-y-3">
                    {players.map(player => player.user && (
                      <li key={player.id} className="flex items-center">
                        <div className="flex-shrink-0">
                          {player.user.avatar_url ? (
                            <img 
                              src={player.user.avatar_url} 
                              alt={player.user.username} 
                              className="h-8 w-8 rounded-full"
                            />
                          ) : (
                            <div className="h-8 w-8 rounded-full bg-gray-500 flex items-center justify-center">
                              <span className="text-white text-sm font-medium">
                                {player.user.username.charAt(0).toUpperCase()}
                              </span>
                            </div>
                          )}
                        </div>
                        <div className="ml-3">
                          <p className="text-sm font-medium text-gray-900 dark:text-white">
                            {player.user.username}
                          </p>
                        </div>
                      </li>
                    ))}
                  </ul>
                </div>
              )}
              
              {players.length === 0 && (
                <p className="text-gray-500 dark:text-gray-500 italic text-sm">
                  No players have joined this campaign yet.
                </p>
              )}
            </div>
          </div>
        </div>
      </div>
    );
  } catch (error) {
    console.error('Error loading campaign:', error);
    return notFound();
  }
}