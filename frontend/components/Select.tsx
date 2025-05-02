import React from 'react';
import { SelectOption } from '../app/api/types';

interface SelectProps extends Omit<React.SelectHTMLAttributes<HTMLSelectElement>, 'onChange'> {
  label: string;
  options: SelectOption[];
  error?: string;
  fullWidth?: boolean;
  onChange?: (value: string) => void;
}

const Select: React.FC<SelectProps> = ({
  label,
  options,
  error,
  fullWidth = true,
  className = '',
  id,
  onChange,
  value,
  ...props
}) => {
  const selectId = id || label.toLowerCase().replace(/\s+/g, '-');
  
  const handleChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    if (onChange) {
      onChange(e.target.value);
    }
  };
  
  return (
    <div className={`mb-4 ${fullWidth ? 'w-full' : ''}`}>
      <label 
        htmlFor={selectId} 
        className="block text-sm font-medium text-gray-700 mb-1"
      >
        {label}
      </label>
      <select
        id={selectId}
        className={`
          px-3 py-2 bg-white border shadow-sm border-gray-300 
          focus:outline-none focus:border-purple-500 focus:ring-purple-500 block w-full 
          rounded-md sm:text-sm focus:ring-1 
          ${error ? 'border-red-500' : ''}
          ${className}
        `}
        onChange={handleChange}
        value={value}
        {...props}
      >
        <option value="" disabled>Select {label}</option>
        {options.map((option) => (
          <option key={option.value} value={option.value}>
            {option.label}
          </option>
        ))}
      </select>
      {error && (
        <p className="mt-1 text-sm text-red-600">{error}</p>
      )}
    </div>
  );
};

export default Select;