import React from "react";
import {FaDragon, FaGithub} from "react-icons/fa";
import Link from "next/link";

export const Footer: React.FC = () => {
  return (
    <footer className="bg-gray-900 text-white py-4">
      <div className="mx-auto max-w-7xl">
        <div className="flex items-center justify-between">
          <div className="flex items-center">
            <FaDragon className="text-purple-500 h-6 w-6" />
            <span className="ml-2 text-lg font-bold">LoreCrafter</span>
          </div>
          <div className="flex justify-between">
            <div>
              <h3 className="text-sm font-semibold uppercase tracking-wider">Navigation</h3>
              <div className="flex flex-col mt-2 space-y-2">
                <Link href="/" className="text-gray-300 hover:text-white block">Home</Link>
                <Link href="/character" className="text-gray-300 hover:text-white block">Characters</Link>
                <Link href="/world" className="text-gray-300 hover:text-white block">Worlds</Link>
                <Link href="/campaign" className="text-gray-300 hover:text-white block">Campaigns</Link>
              </div>
            </div>
            <div className="ml-8">
              <h3 className="text-sm font-semibold uppercase tracking-wider">Resources</h3>
              <div className="flex flex-col mt-2 space-y-2">
                <Link href="/search" className="text-gray-300 hover:text-white block">Search</Link>
                <Link href="/github" className="text-gray-300 hover:text-white block flex">
                  <FaGithub className="h-4 w-4"/>
                  <span className="ml-1">Github</span>
                </Link>
              </div>
            </div>
          </div>
        </div>
        <div className="flex items-center justify-center border-t border-gray-800 pt-8 mt-8 text-sm">
          <p>&copy; {new Date().getFullYear()} LoreCrafter. All rights reserved.</p>
        </div>
      </div>
    </footer>
  )
}