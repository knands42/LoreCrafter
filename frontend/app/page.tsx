import React from 'react';
import Link from 'next/link';
import { FaUser, FaGlobe, FaDungeon, FaPlus } from 'react-icons/fa';

export default function Home() {
  return (
      <div className="space-y-8">
        {/* Hero Section */}
        <section className="text-center py-12 px-4 sm:px-6 lg:px-8 bg-gradient-to-r from-purple-700 to-indigo-800 rounded-lg shadow-xl text-white">
          <h1 className="text-4xl font-extrabold tracking-tight sm:text-5xl lg:text-6xl">
            Welcome to LoreCrafter
          </h1>
          <p className="mt-6 text-xl max-w-3xl mx-auto">
            Create rich backstories for your tabletop role-playing game characters, worlds, and campaigns using AI.
          </p>
          <div className="mt-8 flex justify-center space-x-4">
            <Link 
              href="/characters/create" 
              className="inline-flex items-center px-6 py-3 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-purple-500 hover:bg-purple-600"
            >
              <FaPlus className="mr-2" />
              Create Character
            </Link>
            <Link 
              href="/worlds/create" 
              className="inline-flex items-center px-6 py-3 border border-transparent text-base font-medium rounded-md shadow-sm text-purple-600 bg-white hover:bg-gray-50"
            >
              <FaPlus className="mr-2" />
              Create World
            </Link>
          </div>
        </section>

        {/* Dashboard Sections */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          {/* Characters Section */}
          <section className="bg-white dark:bg-gray-700 shadow rounded-lg overflow-hidden">
            <div className="bg-purple-600 px-4 py-5 sm:px-6 flex items-center justify-between">
              <h2 className="text-lg font-medium text-white flex items-center">
                <FaUser className="mr-2" />
                Characters
              </h2>
              <Link 
                href="/characters" 
                className="text-sm text-purple-100 hover:text-white"
              >
                View All
              </Link>
            </div>
            <div className="px-4 py-5 sm:p-6">
              <p className="text-gray-600 dark:text-gray-300 mb-4">
                Create and manage your character backstories, personalities, and appearances.
              </p>
              <Link 
                href="/characters/create" 
                className="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-purple-600 hover:bg-purple-700"
              >
                <FaPlus className="mr-2" />
                New Character
              </Link>
            </div>
          </section>

          {/* Worlds Section */}
          <section className="bg-white dark:bg-gray-700 shadow rounded-lg overflow-hidden">
            <div className="bg-indigo-600 px-4 py-5 sm:px-6 flex items-center justify-between">
              <h2 className="text-lg font-medium text-white flex items-center">
                <FaGlobe className="mr-2" />
                Worlds
              </h2>
              <Link 
                href="/worlds" 
                className="text-sm text-indigo-100 hover:text-white"
              >
                View All
              </Link>
            </div>
            <div className="px-4 py-5 sm:p-6">
              <p className="text-gray-600 dark:text-gray-300 mb-4">
                Create and manage your world histories, timelines, and settings.
              </p>
              <Link 
                href="/worlds/create" 
                className="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700"
              >
                <FaPlus className="mr-2" />
                New World
              </Link>
            </div>
          </section>

          {/* Campaigns Section */}
          <section className="bg-white dark:bg-gray-700 shadow rounded-lg overflow-hidden">
            <div className="bg-blue-600 px-4 py-5 sm:px-6 flex items-center justify-between">
              <h2 className="text-lg font-medium text-white flex items-center">
                <FaDungeon className="mr-2" />
                Campaigns
              </h2>
              <Link 
                href="/campaigns" 
                className="text-sm text-blue-100 hover:text-white"
              >
                View All
              </Link>
            </div>
            <div className="px-4 py-5 sm:p-6">
              <p className="text-gray-600 dark:text-gray-300 mb-4">
                Create and manage your campaigns, linking them to worlds and characters.
              </p>
              <Link 
                href="/campaigns/create" 
                className="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700"
              >
                <FaPlus className="mr-2" />
                New Campaign
              </Link>
            </div>
          </section>
        </div>

        {/* Search Section */}
        <section className="bg-white dark:bg-gray-700 shadow rounded-lg overflow-hidden">
          <div className="px-4 py-5 sm:p-6">
            <h2 className="text-lg font-medium text-gray-900 dark:text-white mb-4">
              Search Your Content
            </h2>
            <p className="text-gray-600 dark:text-gray-300 mb-4">
              Looking for a specific character or world? Use our search to find it quickly.
            </p>
            <Link 
              href="/search" 
              className="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-gray-600 hover:bg-gray-700"
            >
              Go to Search
            </Link>
          </div>
        </section>
      </div>
  );
}