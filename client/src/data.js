export const sportsData = [
  {
    name: "cricket",
    fixtures: [
      {
        id: "9703206",
        homeId: "202949",
        homeName: "Rajasthan Royals",
        awayId: "162053",
        awayName: "Mumbai Indians",
      },
    ],
  },
  {
    name: "volleyball",
    fixtures: [
      {
        id: "9879535",
        homeId: "210549",
        homeName: "Sporting CP Women",
        awayId: "307260",
        awayName: "FC Porto Women",
      },
    ],
  },
];

export const bets = {
  "1x2": [
    { id: "658666625", odds: "2.50", name: "1", header: "" },
    { id: "658666628", odds: "1.53", name: "2", header: "" },
  ],
  underOverCricket: [
    { id: "658772390", odds: "1.83", name: "186.5", header: "Over" },
    { id: "658772394", odds: "1.83", name: "186.5", header: "Under" },
  ],
  underOverVolley: [
    {
      id: "670136314",
      odds: "1.83",
      name: "Total",
      header: "1",
      handicap: "O 45.5",
    },
    {
      id: "670136315",
      odds: "1.83",
      name: "Total",
      header: "2",
      handicap: "U 45.5",
    },
  ],
  correctScore: [
    { Id: "670136282", Odds: "3.10", Name: "3-0", Header: "1" },
    { Id: "670136286", Odds: "4.00", Name: "3-1", Header: "1" },
    { Id: "670136288", Odds: "5.00", Name: "3-2", Header: "1" },
    { Id: "670136294", Odds: "8.00", Name: "3-0", Header: "2" },
    { Id: "670136301", Odds: "7.50", Name: "3-1", Header: "2" },
    { Id: "670136306", Odds: "7.00", Name: "3-2", Header: "2" },
  ],
};
