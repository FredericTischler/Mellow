'use client';

import { useRef, useState, useEffect } from 'react';

export default function FileInput({
  label = 'Choose a file',
  onChange,
  id,
  name,
  setValue,
  usePreview = true,
  isMini = false,
}) {
  const inputRef = useRef(null);
  const [fileName, setFileName] = useState('');
  const [preview, setPreview] = useState(null);

  const handleClick = () => {
    inputRef.current?.click();
  };

  const handleFileChange = (e) => {
    const file = e.target.files[0];
    if (!file) return;

    setFileName(file.name);
    if (file.type.startsWith('image/')) {
      setPreview(URL.createObjectURL(file));
    } else {
      setPreview(null);
    }

    // Injection dans react-hook-form
    setValue(name, file);

    // Callback optionnel
    onChange?.(file);
  };

  useEffect(() => {
    return () => {
      if (preview) URL.revokeObjectURL(preview);
    };
  }, [preview]);

  return (
    <div className="flex flex-col gap-2 items-center">
      <div className="flex justify-start items-center gap-2 w-full">
        <button
          type="button"
          onClick={handleClick}
          className={`inline-flex items-center gap-2 px-4 py-2  text-lavender-5 border border-lavender-5 rounded-2xl cursor-pointer transition w-fit ${isMini && fileName?'bg-lavender-1 text-lavender-2':' bg-transparent  hover:bg-lavender-6'}`}
        >
          {label}
        </button>

        <input
          type="file"
          id={id || name}
          name={name}
          accept="image/*"
          ref={inputRef}
          onChange={handleFileChange}
          className="hidden"
        />
        {fileName && !isMini && <span className="text-sm">{fileName}</span>}
      </div>
      {preview && usePreview && (
        <img
          src={preview}
          alt="Preview"
          className="w-24 h-24 object-cover rounded-full border border-lavender-3 shadow-(--box-shadow)"
        />
      )}
    </div>
  );
}
