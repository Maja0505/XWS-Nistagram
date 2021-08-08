export const renderer = ({ story, action, isPaused, config }) => {
	return <div>Hello!</div>;
};

export const tester = (story) => {
	return {
		// Use this renderer only when the story type is video
		condition: story.type === 'video',
		priority: 3,
	};
};