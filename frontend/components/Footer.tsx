import React from 'react';
import Link from 'next/link';
import { FaDragon, FaGithub } from 'react-icons/fa';

const Footer: React.FC = () => {
  return (
    <footer className="bg-gray-900 text-white py-8">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex flex-col md:flex-row justify-between items-center">
          <div className="flex items-center mb-4 md:mb-0">
            <FaDragon className="h-6 w-6 text-purple-500" />
            <span className="ml-2 text-lg font-bold">LoreCrafter</span>
          </div>
          
          <div className="flex flex-col md:flex-row space-y-4 md:space-y-0 md:space-x-8">
            <div>
              <h3 className="text-sm font-semibold uppercase tracking-wider">Navigation</h3>
              <div className="mt-2 space-y-2">
                <Link href="/public" className="text-gray-300 hover:text-white block">
                  Home
                </Link>
                <Link href="/characters" className="text-gray-300 hover:text-white block">
                  Characters
                </Link>
                <Link href="/worlds" className="text-gray-300 hover:text-white block">
                  Worlds
                </Link>
                <Link href="/campaigns" className="text-gray-300 hover:text-white block">
                  Campaigns
                </Link>
              </div>
            </div>
            
            <div>
              <h3 className="text-sm font-semibold uppercase tracking-wider">Resources</h3>
              <div className="mt-2 space-y-2">
                <Link href="/search" className="text-gray-300 hover:text-white block">
                  Search
                </Link>
                <a 
                  href="https://github.com/yourusername/lorecrafter" 
                  target="_blank" 
                  rel="noopener noreferrer"
                  className="text-gray-300 hover:text-white flex items-center"
                >
                  <FaGithub className="mr-1" />
                  GitHub
                </a>
              </div>
            </div>
          </div>
        </div>
        
        <div className="mt-8 border-t border-gray-800 pt-8 text-center text-sm text-gray-400">
          <p>&copy; {new Date().getFullYear()} LoreCrafter. All rights reserved.</p>
        </div>
      </div>
    </footer>
  );
};

export default Footer;