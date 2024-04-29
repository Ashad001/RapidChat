import React, { useState } from 'react';
import EmojiPicker from 'emoji-picker-react'; // Ensure this is correctly imported

function SelectEmoji() {
  const [selectedEmoji, setSelectedEmoji] = useState(null);

  function handleEmojiSelect(emoji) {
    setSelectedEmoji(emoji);
  }

  return (
    <div>
      <h1>Selected Emoji: {selectedEmoji}</h1>
      <EmojiPicker
      onEmojiSelect={handleEmojiSelect}
      rows={4}
      perRow={8}
      emojiSize={32}
      pickerStyle={{ position: 'absolute', bottom: '20px', right: '20px' }}
    />
    </div>
  );
}

export default SelectEmoji;
