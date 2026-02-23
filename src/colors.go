package main

func split(faces [6][9]int) [3][3][3][6]int {
	noColor := LIGHT_BLACK
	return [3][3][3][6]int{
		{ //1
			{ //1
				{ //1
					noColor, faces[LEFT][0], faces[BACK][0], noColor, noColor, faces[BOTTOM][0],
				},
				{ //2
					noColor, faces[LEFT][1], noColor, noColor, noColor, faces[BOTTOM][1],
				},
				{ //3
					faces[FRONT][0], faces[LEFT][2], noColor, noColor, noColor, faces[BOTTOM][2],
				},
			},
			{ //2
				{ //1
					noColor, faces[LEFT][3], faces[BACK][1], noColor, noColor, noColor,
				},
				{ //2
					noColor, faces[LEFT][4], noColor, noColor, noColor, noColor,
				},
				{ //3
					faces[FRONT][1], faces[LEFT][5], noColor, noColor, noColor, noColor,
				},
			},
			{ //3
				{ //1
					noColor, faces[LEFT][6], faces[BACK][2], noColor, faces[TOP][0], noColor,
				},
				{ //2
					noColor, faces[LEFT][7], noColor, noColor, faces[TOP][1], noColor,
				},
				{ //3
					faces[FRONT][2], faces[LEFT][8], noColor, noColor, faces[TOP][2], noColor,
				},
			},
		},
		{ //2
			{ //1
				{ //1
					noColor, noColor, faces[BACK][3], noColor, noColor, faces[BOTTOM][3],
				},
				{ //2
					noColor, noColor, noColor, noColor, noColor, faces[BOTTOM][4],
				},
				{ //3
					faces[FRONT][3], noColor, noColor, noColor, noColor, faces[BOTTOM][5],
				},
			},
			{ //2
				{ //1
					noColor, noColor, faces[BACK][4], noColor, noColor, noColor,
				},
				{ //2
					noColor, noColor, noColor, noColor, noColor, noColor,
				},
				{ //3
					faces[FRONT][4], noColor, noColor, noColor, noColor, noColor,
				},
			},
			{ //3
				{ //1
					noColor, noColor, faces[BACK][5], noColor, faces[TOP][3], noColor,
				},
				{ //2
					noColor, noColor, noColor, noColor, faces[TOP][4], noColor,
				},
				{ //3
					faces[FRONT][5], noColor, noColor, noColor, faces[TOP][5], noColor,
				},
			},
		},
		{ //3
			{ //1
				{ //1
					noColor, noColor, faces[BACK][6], faces[RIGHT][0], noColor, faces[BOTTOM][6],
				},
				{ //2
					noColor, noColor, noColor, faces[RIGHT][1], noColor, faces[BOTTOM][7],
				},
				{ //3
					faces[FRONT][6], noColor, noColor, faces[RIGHT][2], noColor, faces[BOTTOM][8],
				},
			},
			{ //2
				{ //1
					noColor, noColor, faces[BACK][7], faces[RIGHT][3], noColor, noColor,
				},
				{ //2
					noColor, noColor, noColor, faces[RIGHT][4], noColor, noColor,
				},
				{ //3
					faces[FRONT][7], noColor, noColor, faces[RIGHT][5], noColor, noColor,
				},
			},
			{ //3
				{ //1
					noColor, noColor, faces[BACK][8], faces[RIGHT][6], faces[TOP][6], noColor,
				},
				{ //2
					noColor, noColor, noColor, faces[RIGHT][7], faces[TOP][7], noColor,
				},
				{ //3
					faces[FRONT][8], noColor, noColor, faces[RIGHT][8], faces[TOP][8], noColor,
				},
			},
		},
	}
}
