import React from "react";
import Link from "next/link";
import {FaDragon, FaSignInAlt} from "react-icons/fa";

export const Navbar: React.FC = () => {
  return (
    <nav className="bg-gray-900 text-white shadow-lg">
      <div className="flex justify-between max-w-7xl mx-auto">
        <div className="flex items-center h-16">
          <Link href="/" className="flex items-center">
            <FaDragon className="h-8 w-8 text-purple-500"/>
            <span className="ml-2 text-lg font-bold">LoreCrafter</span>
          </Link>
        </div>

        <div className="flex items-center h-16">
          <Link href="/auth/signin" className="flex items-center">
            <FaSignInAlt className="h-8 w-8 text-purple-500"/>
            <span className="ml-2 text-lg">Sign In</span>
          </Link>
        </div>
      </div>
    </nav>
  )
}