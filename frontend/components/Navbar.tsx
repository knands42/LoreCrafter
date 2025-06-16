'use client';

import React from "react";
import Link from "next/link";
import {FaDragon, FaDungeon, FaGlobe, FaSignInAlt, FaSignOutAlt, FaUser} from "react-icons/fa";
import {useAuth} from "@/context/auth";
import {ThemeToggle} from "./ThemeToggle";

export const Navbar: React.FC = () => {
  const { user, loading, logout } = useAuth();

  return (
    <nav className="bg-gray-900 text-white shadow-lg">
      <div className="flex justify-between max-w-7xl mx-auto px-4">
        <div className="flex items-center h-16">
          <Link href="/" className="flex items-center">
            <FaDragon className="h-8 w-8 text-purple-500"/>
            <span className="ml-2 text-lg font-bold">LoreCrafter</span>
          </Link>

          {user && (
            <div className="hidden md:flex ml-8 space-x-6">
              <Link href="/campaign" className="flex items-center text-gray-300 hover:text-white">
                <FaDungeon className="h-5 w-5 text-purple-500 mr-1"/>
                <span>Campaigns</span>
              </Link>
              <Link href="/character" className="flex items-center text-gray-300 hover:text-white">
                <FaUser className="h-5 w-5 text-purple-500 mr-1"/>
                <span>Characters</span>
              </Link>
              <Link href="/world" className="flex items-center text-gray-300 hover:text-white">
                <FaGlobe className="h-5 w-5 text-purple-500 mr-1"/>
                <span>Worlds</span>
              </Link>
            </div>
          )}
        </div>

        <div className="flex items-center h-16 space-x-4">
          <ThemeToggle />

          {loading ? (
            <div className="h-8 w-8 rounded-full bg-gray-700 animate-pulse"></div>
          ) : user ? (
            <div className="flex items-center space-x-4">
              <Link href="/profile" className="flex items-center text-gray-300 hover:text-white">
                <div className="relative">
                  {user.avatar_url ? (
                    <img 
                      src={user.avatar_url} 
                      alt={user.username} 
                      className="h-8 w-8 rounded-full"
                    />
                  ) : (
                    <div className="h-8 w-8 rounded-full bg-purple-600 flex items-center justify-center">
                      <span className="text-white text-sm font-medium">
                        {user.username.charAt(0).toUpperCase()}
                      </span>
                    </div>
                  )}
                </div>
                <span className="ml-2 hidden md:block">{user.username}</span>
              </Link>

              <button 
                onClick={() => logout()} 
                className="flex items-center text-gray-300 hover:text-white"
              >
                <FaSignOutAlt className="h-5 w-5 text-purple-500"/>
                <span className="ml-1 hidden md:block">Sign Out</span>
              </button>
            </div>
          ) : (
            <Link href="/auth/signin" className="flex items-center">
              <FaSignInAlt className="h-5 w-5 text-purple-500"/>
              <span className="ml-2">Sign In</span>
            </Link>
          )}
        </div>
      </div>
    </nav>
  );
}
