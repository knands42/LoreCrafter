import React from 'react';
import Link from 'next/link';
import { FaDragon, FaSearch, FaUser, FaGlobe, FaDungeon } from 'react-icons/fa';

const Navbar: React.FC = () => {
  return (
    <nav className="bg-gray-900 text-white shadow-lg">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between h-16">
          <div className="flex items-center">
            <Link href="/" className="flex items-center">
              <FaDragon className="h-8 w-8 text-purple-500" />
              <span className="ml-2 text-xl font-bold">LoreCrafter</span>
            </Link>
          </div>
          
          <div className="flex items-center space-x-4">
            <Link 
              href="/characters" 
              className="px-3 py-2 rounded-md text-sm font-medium hover:bg-gray-800 flex items-center"
            >
              <FaUser className="mr-1" />
              Characters
            </Link>
            
            <Link 
              href="/worlds" 
              className="px-3 py-2 rounded-md text-sm font-medium hover:bg-gray-800 flex items-center"
            >
              <FaGlobe className="mr-1" />
              Worlds
            </Link>
            
            <Link 
              href="/campaigns" 
              className="px-3 py-2 rounded-md text-sm font-medium hover:bg-gray-800 flex items-center"
            >
              <FaDungeon className="mr-1" />
              Campaigns
            </Link>
            
            <Link 
              href="/search" 
              className="px-3 py-2 rounded-md text-sm font-medium hover:bg-gray-800 flex items-center"
            >
              <FaSearch className="mr-1" />
              Search
            </Link>
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;