module.exports = {
	branches: ['main'],
	plugins: [
		'@semantic-release/commit-analyzer',
		'@semantic-release/release-notes-generator',
		['@semantic-release/changelog', { changelogFile: 'CHANGELOG.md' }],
		[
			'@semantic-release/npm',
			{
				npmPublish: true,
				pkgRoot: '.',
				registry: 'https://registry.npmjs.org/',
			},
		],
		'@semantic-release/github',
		[
			'@semantic-release/git',
			{
				assets: ['CHANGELOG.md'],
				message:
					'chore(release): set `package.json` to ${nextRelease.version} [skip ci]\n\n${nextRelease.notes}',
			},
		],
	],
};
